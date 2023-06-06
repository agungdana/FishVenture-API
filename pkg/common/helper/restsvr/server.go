package restsvr

import (
	"net/http"

	"github.com/e-fish/api/pkg/common/helper/ctxutil"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func NewRoute() *gin.Engine {
	r := gin.Default()

	r.Use(ctxutil.Authentication())

	return r
}

func Run(r *gin.Engine) {

	r.GET("/status", func(ctx *gin.Context) {
		res := new(HttpResponse)
		res.Add("running", nil)
		ResponsJson(ctx, res)
	})

	// if err := r.Run(conf.AppsAddress + ":" + conf.AppsPort); err != nil {
	// 	logger.Fatal("failed to run serv: %v", err)
	// }
	r.Run()
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return upgrader.Upgrade(w, r, nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
