package ctxutil

import (
	"strings"

	"github.com/e-fish/api/pkg/common/infra/token"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = NewRequest(ctx)
		header := c.GetHeader("Authorization")
		if header == "" {
			c.Request = c.Request.WithContext(ctx)
			c.Next()
			return
		}
		headers := strings.Split(header, " ")
		if len(headers) == 0 || len(headers) != 2 {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "authentication doesn't exist",
			})
		}
		if headers[0] != "Bearer" {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "authentication not supported",
			})
			return
		}

		tokenMaker, _ := token.NewTokenMaker(token.SecretKey)
		payload := token.Payload{}
		err := tokenMaker.VerifyToken(headers[1], &payload)
		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "authentication not match :" + err.Error(),
			})
			return
		}

		ctx = SetUserPayload(ctx, payload.UserID, payload.PondID, payload.AppType, payload.UserRole...)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		path := c.Request.URL.Path

		if !CanAccess(ctx, path) {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "access denied",
			})
		}

		c.Next()
	}
}
