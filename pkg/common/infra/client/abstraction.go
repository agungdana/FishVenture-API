package client

import (
	"context"

	"github.com/google/uuid"
)

type PubSub interface {
	Subscribe(ctx context.Context, topic string, f NatsClientMesage) error
	Publish(ctx context.Context, data any, topic string, target uuid.UUID)
	Close(ctx context.Context, id uuid.UUID, reqID uuid.UUID) error
}
