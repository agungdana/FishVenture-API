package redis

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/event"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var (
	subs = &Subscribtions{
		mut:           sync.Mutex{},
		subscribtions: make(map[uuid.UUID]*subscribtion),
	}

	redisSubOnce sync.Once
)

type Subscribtions struct {
	mut           sync.Mutex
	subscribtions map[uuid.UUID]*subscribtion
}

type subscribtion struct {
	pubSub            *redis.PubSub
	callbackMessaging map[uuid.UUID]event.ClientMessaging
	counter           int
}

func (s *Subscribtions) getID(ctx context.Context) uuid.UUID {
	var (
		userID, _ = ctxutil.GetUserID(ctx)
		pondID, _ = ctxutil.GetPondID(ctx)
	)

	if pondID != uuid.Nil {
		return pondID
	}

	if userID != uuid.Nil {
		return userID
	}

	return uuid.New()
}

func (s *Subscribtions) add(ctx context.Context, redisSubs *redis.PubSub, msg event.ClientMessaging) uuid.UUID {
	var (
		requestID, _ = ctxutil.GetRequestID(ctx)
		id           = s.getID(ctx)
	)

	if id == uuid.Nil {
		id = uuid.New()
	}

	s.mut.Lock()
	defer s.mut.Unlock()
	_, ok := s.subscribtions[id]
	if !ok {
		s.subscribtions[id] = &subscribtion{
			pubSub:            redisSubs,
			callbackMessaging: make(map[uuid.UUID]event.ClientMessaging),
		}
		go s.ReceiveMessage(ctx, id)
	}

	s.subscribtions[id].counter += 1
	s.subscribtions[id].callbackMessaging[requestID] = msg

	return id
}

func (s *Subscribtions) close(ctx context.Context, id uuid.UUID) error {
	var (
		requestID, _ = ctxutil.GetRequestID(ctx)
	)

	s.mut.Lock()
	defer s.mut.Unlock()
	val, ok := s.subscribtions[id]
	if !ok {
		return nil
	}

	if val.counter < 1 {
		err := val.pubSub.Close()
		if err != nil {
			return err
		}

		delete(s.subscribtions, id)
		return nil
	}

	s.subscribtions[id].counter -= 1
	delete(s.subscribtions[id].callbackMessaging, requestID)

	return nil
}

func newRedisEventPubSub(conf config.RedisConfig, client *redis.Client) event.EventPubsub {
	redisSubOnce.Do(func() {
		subs = &Subscribtions{
			subscribtions: make(map[uuid.UUID]*subscribtion),
		}
	})

	return &redisEventPubSub{
		conf:          conf,
		redis:         client,
		subscribtions: subs,
	}
}

type redisEventPubSub struct {
	conf          config.RedisConfig
	redis         *redis.Client
	subscribtions *Subscribtions
}

// Close implements EventPubsub.
func (r *redisEventPubSub) Close(ctx context.Context, id uuid.UUID) error {

	if id == uuid.Nil {
		id = r.subscribtions.getID(ctx)
	}

	return r.subscribtions.close(ctx, id)
}

// Publish implements EventPubsub.
func (r *redisEventPubSub) Publish(ctx context.Context, msg event.ClientMessages) (err error) {

	ctx = ctxutil.CopyCtxWithoutTimeout(ctx)
	if msg.CtxMap == nil {
		msg.CtxMap = ctxutil.ToMap(ctx)
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return werror.Error{
			Code:    "FailedMarshalMessage",
			Message: "failed json.Marshal message",
			Details: map[string]any{
				"err": err,
			},
		}
	}

	err = r.redis.Publish(ctx, msg.Topic, data).Err()

	if err != nil {
		err = werror.Error{
			Code:    "FailedPublishMessage",
			Message: "failed publish message to topic",
			Details: map[string]any{"error": err},
		}
	}

	return err
}

// Subscribe implements EventPubsub.
func (r *redisEventPubSub) Subscribe(ctx context.Context, topic string, msg event.ClientMessaging) (uuid.UUID, error) {

	redisSubs := r.redis.PSubscribe(ctx, topic)
	if redisSubs == nil {
		return uuid.Nil, werror.Error{
			Code:    "FailedSubscibeChanel",
			Message: "failed added new subscribtion",
			Details: map[string]any{"channel": topic},
		}
	}

	id := r.subscribtions.add(ctx, redisSubs, msg)

	return id, nil
}

func (r *Subscribtions) ReceiveMessage(ctx context.Context, id uuid.UUID) error {

	subscribtion, ok := r.subscribtions[id]
	if !ok {
		err := werror.Error{
			Code:    "SubscribtionNotFound",
			Message: "subscribtion not found",
		}
		logger.ErrorWithContext(ctx, "##Failed receive message redis err: %s", err)
		return err
	}

	for {
		redisMessage, err := subscribtion.pubSub.Receive(ctx)
		if err != nil {
			return werror.Error{
				Code:    "FailedReceiveMessage",
				Message: "failed receive message",
				Details: map[string]any{"error": err, "msg": redisMessage},
			}
		}

		switch rm := redisMessage.(type) {
		case *redis.Message:
			for _, send := range subscribtion.callbackMessaging {

				cMsg := event.ClientMessages{}

				err := json.Unmarshal([]byte(rm.Payload), &cMsg)
				if err != nil {
					logger.ErrorWithContext(ctx, "##Failed json.Unmarshal message redis err: [%s]", err)
				}

				send(&event.ClientMessages{
					Topic:  cMsg.Topic,
					Action: cMsg.Action,
					CtxMap: cMsg.CtxMap,
					Data:   cMsg.Data,
				})
			}
		}
	}
}
