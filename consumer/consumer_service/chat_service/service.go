package chatservice

import (
	"context"

	chatconfig "github.com/e-fish/api/chat_http/chat_config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/e-fish/api/pkg/common/infra/event"
	"github.com/e-fish/api/pkg/domain/chat"
)

type ChatService struct {
	chatConf chatconfig.ChatConfig
	event    event.EventPubsub
	repo     chat.Repo
}

func NewService(event event.EventPubsub) ChatService {
	var (
		chatConf = chatconfig.GetConfig()
		service  = ChatService{
			chatConf: *chatConf,
		}
	)

	chatRepo, err := chat.NewRepo(chatConf.ChatDBConfig)
	if err != nil {
		logger.Fatal("failed to create a new repo tenant, can't create consumer chat service err: %v", err)
	}

	service.repo = chatRepo
	service.event = event

	return service
}

func (s *ChatService) PublishChatToAll(ctx context.Context, req ReadChatRequest) error {
	query := s.repo.NewQuery()

	chat, err := query.ReadChatByID(ctx, req.ChatID)
	if err != nil {
		err = werror.Error{
			Code:    "ConsumerOrderServicePublishOrderToAll",
			Message: "failed query.ReadOrderByID on [consumer.OrderService.PublishOrderToAll]",
			Details: map[string]any{
				"error": err,
			},
		}
		logger.ErrorWithContext(ctx, "failed query.ReadOrderByID on [consumer.OrderService.PublishOrderToAll] ")
		return err
	}

	cs := event.GetTopicUser(string(event.WS_REPLY), chat.UserID, s.chatConf.AppConfig.Debug)
	pond := event.GetTopicUser(string(event.WS_REPLY), chat.PondID, s.chatConf.AppConfig.Debug)
	wsRess := restsvr.WsResponse{
		Key:   "chat",
		Value: chat,
	}

	s.event.Publish(ctx, event.ClientMessages{
		Topic:  cs,
		Action: "chat",
		Data:   wsRess,
	})

	s.event.Publish(ctx, event.ClientMessages{
		Topic:  pond,
		Action: "chat",
		Data:   wsRess,
	})

	return nil
}

func (s *ChatService) PublishChatItemToAll(ctx context.Context, req ReadChatRequest) error {
	query := s.repo.NewQuery()

	chat, err := query.ReadChatItemsByID(ctx, req.ChatID)
	if err != nil {
		err = werror.Error{
			Code:    "ConsumerOrderServicePublishOrderToAll",
			Message: "failed query.ReadOrderByID on [consumer.OrderService.PublishOrderToAll]",
			Details: map[string]any{
				"error": err,
			},
		}
		logger.ErrorWithContext(ctx, "failed query.ReadOrderByID on [consumer.OrderService.PublishOrderToAll] ")
		return err
	}

	cs := event.GetTopicUser(string(event.WS_REPLY), chat.SenderID, s.chatConf.AppConfig.Debug)
	pond := event.GetTopicUser(string(event.WS_REPLY), chat.ReceiverID, s.chatConf.AppConfig.Debug)
	wsRess := restsvr.WsResponse{
		Key:   "chat_item",
		Value: chat,
	}

	s.event.Publish(ctx, event.ClientMessages{
		Topic:  cs,
		Action: "chat_item",
		Data:   wsRess,
	})

	s.event.Publish(ctx, event.ClientMessages{
		Topic:  pond,
		Action: "chat_item",
		Data:   wsRess,
	})

	return nil
}
