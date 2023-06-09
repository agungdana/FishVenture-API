package authservice

import (
	"context"

	authconfig "github.com/e-fish/api/auth_http/auth_config"
)

type Service struct {
	conf authconfig.AuthConfig
}

func NewService(conf authconfig.AuthConfig) Service {
	var (
		ctx     = context.Background()
		service = Service{
			conf: conf,
		}
	)

	go service.RegisterPermissionAccess(ctx)

	return service
}

func (s *Service) RegisterPermissionAccess(ctx context.Context) error {

	return nil
}
