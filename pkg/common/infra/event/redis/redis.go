package redis

import (
	"fmt"
	"sync"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/infra/event"
	"github.com/redis/go-redis/v9"
)

var (
	cl         *redisClientPool
	clientOnce sync.Once
)

type redisClientPool struct {
	clients map[string]*redis.Client
}

func getRedisClientPool() *redisClientPool {
	if cl == nil {
		clientOnce.Do(func() {
			cl = &redisClientPool{
				clients: map[string]*redis.Client{},
			}
		})
	}

	return cl
}

func NewRedisClient(conf config.RedisConfig) (event.Event, error) {
	rClient, err := createConnectionRedis(conf)
	if err != nil {
		return nil, err
	}

	client := redisClient{
		conf:   conf,
		client: rClient,
	}

	return &client, nil
}

type redisClient struct {
	conf   config.RedisConfig
	client *redis.Client
}

// NewEventPubsub implements
func (r *redisClient) NewEventPubsub() event.EventPubsub {
	return newRedisEventPubSub(r.conf, r.client)
}

// NewEventStore implements
func (r *redisClient) NewEventStore() event.EventStore {
	return newRedisEventStore(r.conf, r.client)
}

func createConnectionRedis(conf config.RedisConfig) (*redis.Client, error) {

	redisAddr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	cl, ok := getRedisClientPool().clients[redisAddr]
	if ok {
		return cl, nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: conf.Password,
		DB:       conf.DB,
	})

	if client == nil {
		return nil, ErrFailedCreateRedisConnection
	}

	cl = client

	return client, nil
}
