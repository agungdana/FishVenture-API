package event

import (
	"context"

	"github.com/google/uuid"
)

type Event interface {
	NewEventPubsub() EventPubsub
	NewEventStore() EventStore
}

type EventStore interface {
	Set(ctx context.Context, param Payload) (err error)
	Get(ctx context.Context, key string, data any) error
}

type EventPubsub interface {
	Subscribe(ctx context.Context, topic string, msg ClientMessaging) (uuid.UUID, error)
	Publish(ctx context.Context, msg ClientMessages) (err error)
	Close(ctx context.Context, id uuid.UUID) error
}
