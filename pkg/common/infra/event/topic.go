package event

import (
	"context"
	"fmt"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/google/uuid"
)

func GetTopic(ctx context.Context, topic string, isDebug string) string {
	var (
		userID, userOk = ctxutil.GetUserID(ctx)
		pondID, pondOk = ctxutil.GetPondID(ctx)
	)

	if isDebug == "true" {
		topic = fmt.Sprintf("%s.%s", "DEV", topic)
	}

	if pondOk {
		return fmt.Sprintf("%s.%s", topic, pondID)
	}

	if userOk {
		return fmt.Sprintf("%s.%s", topic, userID)
	}

	return fmt.Sprintf("%s.%s", topic, "*")
}

func GetTopicUser(topic string, target uuid.UUID, isDebug string) string {

	if isDebug == "true" {
		topic = fmt.Sprintf("%s.%s", "DEV", topic)
	}

	return fmt.Sprintf("%s.%s", topic, target)
}
