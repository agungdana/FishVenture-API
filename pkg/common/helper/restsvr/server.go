package restsvr

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	restSvr *RestSvr
	once    sync.Once
)

type RestSvr struct {
	conf config.AppConfig
	gin  *gin.Engine
}

func getGinEngine() *RestSvr {
	if restSvr == nil {
		once.Do(func() {
			r := gin.Default()

			r.Use(CORSMiddleware(), Timeout(), ctxutil.Authentication())

			restSvr = &RestSvr{
				gin: r,
			}
		})
	}

	return restSvr
}

func Timeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		if c.Request.Method == "POST" && !strings.Contains(c.Request.URL.Path, "upload-file") {
			ctx, cancel := context.WithTimeout(ctx, time.Second*3)
			defer cancel()

			c.Request = c.Request.WithContext(ctx)
			finished := make(chan struct{})
			go func() {
				c.Next()

				finished <- struct{}{}
			}()

			select {
			case <-finished:
				c.Next()
				return
			case <-ctx.Done():
				ResponsJson(c, &HttpResponse{
					Status:  StatusFailed,
					Message: "connection timeout",
					Error: []ErrorResponse{
						{
							Code:    "RequestTimeout",
							Message: "connection timeout",
							Details: map[string]any{"request to longger": "request > 3s"},
						},
					},
				})

				return
			}
		}
	}
}

func NewRoute(conf config.AppConfig) {
	getGinEngine().conf = conf
}

func GetGinRoute() *gin.Engine {
	return restSvr.gin
}

func Run() {
	rs := getGinEngine()

	rs.gin.GET("/status", func(ctx *gin.Context) {
		res := new(HttpResponse)
		res.Add("running", nil)
		ResponsJson(ctx, res)
	})

	if rs.conf.Host == "" {
		rs.conf.Host = "localhost"
	}
	if rs.conf.Port == "" {
		rs.conf.Port = "8080"
	}

	if err := rs.gin.Run(rs.conf.Host + ":" + rs.conf.Port); err != nil {
		logger.Fatal("failed to run serv: %v", err)
	}

}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return upgrader.Upgrade(w, r, nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
