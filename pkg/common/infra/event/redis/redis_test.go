package redis_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/infra/event"
	"github.com/e-fish/api/pkg/common/infra/event/redis"
	"github.com/stretchr/testify/assert"
)

func TestRedisPubSubEvent(t *testing.T) {
	var (
		ctx   = context.Background()
		topic = "DEV.reply.ab7454b4-a024-47ce-a315-f6ffe71054b5"
	)

	cl, err := redis.NewRedisClient(config.RedisConfig{
		Host:        "localhost",
		Port:        "6379",
		Password:    "",
		DB:          0,
		ExpiredTime: 0,
	})

	assert.NoError(t, err)

	fmt.Printf("topic: %v\n", topic)
	pubsub := cl.NewEventPubsub()
	_, err = pubsub.Subscribe(context.Background(), topic, func(m *event.ClientMessages) {
		fmt.Printf("m.Data: %v\n", m.Data)
	})

	assert.NoError(t, err)

	err = pubsub.Publish(ctx, event.ClientMessages{
		Topic: topic,
		Data:  []byte("coba aja"),
	})

	assert.NoError(t, err)

	time.Sleep(time.Second * 20)

}

type sample struct {
	Data string
}

func TestRedisEventStore(t *testing.T) {
	var (
		ctx  = context.Background()
		key  = "coba-aja"
		data = &sample{
			Data: "ini datanya",
		}
	)

	logger.SetupLogger("true")

	cl, err := redis.NewRedisClient(config.RedisConfig{
		Host:        "localhost",
		Port:        "6379",
		Password:    "",
		DB:          0,
		ExpiredTime: 0,
	})

	assert.NoError(t, err)

	fmt.Printf("key: %v\n", key)
	store := cl.NewEventStore()
	err = store.Set(ctx, event.Payload{
		Key:         key,
		Data:        data,
		ExpiredTime: time.Second * 5,
	})

	assert.NoError(t, err)

	err = store.Get(ctx, key, &data)

	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	data = nil

	time.Sleep(time.Second * 10)

	err = store.Get(ctx, key, &data)
	assert.Error(t, err)
	assert.Empty(t, data)
}

func Test_PublishMessageToTopic(t *testing.T) {

	var (
		ctx = context.Background()
		// topic = "DEV.reply.ab7454b4-a024-47ce-a315-f6ffe71054b5"
		topic = "DEV.reply.41db9e72-a3a2-4caf-aa59-f7e4af79a535"
	)

	cl, err := redis.NewRedisClient(config.RedisConfig{
		Host:        "62.72.31.64",
		Port:        "6379",
		Password:    "",
		DB:          0,
		ExpiredTime: 0,
	})

	assert.NoError(t, err)

	sampleData := sample{
		Data: "ini adalah sebuah data",
	}

	pubsub := cl.NewEventPubsub()

	assert.NoError(t, err)

	err = pubsub.Publish(ctx, event.ClientMessages{
		Topic: topic,
		Data:  sampleData,
	})

	assert.NoError(t, err)

}
