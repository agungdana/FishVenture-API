package authhttp

import (
	authconfig "github.com/e-fish/api/auth_http/auth_config"
	authhandler "github.com/e-fish/api/auth_http/auth_handler"
	authservice "github.com/e-fish/api/auth_http/auth_service"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/gin-gonic/gin"
)

type route struct {
	conf authconfig.AuthConfig
	gin  *gin.Engine
}

func newRoute(ro route) {
	ginEngine := ro.gin

	service := authservice.NewService(ro.conf)
	handler := authhandler.Handler{
		Conf:    ro.conf,
		Service: service,
	}

	ginEngine.POST("/register", handler.CreateUser)
	ginEngine.POST("/update-user", ctxutil.Authorization())
	ginEngine.POST("/delete-user", ctxutil.Authorization())
	ginEngine.POST("/add-user-permission", ctxutil.Authorization())
	ginEngine.POST("/delete-user-permission/id", ctxutil.Authorization())
	ginEngine.POST("/add-user-role", ctxutil.Authorization())
	ginEngine.POST("/delete-user-role/id", ctxutil.Authorization())

	ginEngine.POST("/login", handler.Login)
	ginEngine.POST("/login-by-google", handler.LoginByGoogle)
	ginEngine.GET("/profile", ctxutil.Authorization(), handler.Profile)
}
