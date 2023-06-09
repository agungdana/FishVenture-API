package restsvr_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/e-fish/api/pkg/common/helper/config"
	"github.com/e-fish/api/pkg/common/helper/logger"
	"github.com/e-fish/api/pkg/common/helper/restsvr"
	"github.com/gin-gonic/gin"
)

func TestRestSvrNeRoute(t *testing.T) {
	conf := config.AppConfig{
		Name:    "palen-id",
		Address: "localhost",
		Port:    "8081",
		Debug:   "true",
	}

	logger.SetupLogger(conf.Debug)
	restsvr.NewRoute(conf)

	Route(t)

	restsvr.Run()
}

func Route(t *testing.T) {
	r := restsvr.GetGinRoute()
	r.GET("/test", func(ctx *gin.Context) {
		var (
			res = new(restsvr.HttpResponse)
		)

		defer restsvr.ResponsJson(ctx, res)

		res.Add(map[string]any{"name": "tio", "age": 20}, nil)
	})

	r.POST("/test", func(ctx *gin.Context) {
		var (
			res = new(restsvr.HttpResponse)
			req = map[string]any{}
		)

		defer restsvr.ResponsJson(ctx, res)

		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			res.Add(nil, err)
			return
		}

		res.Add(req, nil)
	})

	r.POST("/test-with-timeout", func(ctx *gin.Context) {
		var (
			res = new(restsvr.HttpResponse)
		)

		defer restsvr.ResponsJson(ctx, res)

		time.Sleep(time.Second * 5)

		res.Add(map[string]any{"name": "tio", "age": 20}, nil)
	})
}

func TestGet(t *testing.T) {
	data := map[string]any{}
	data["nama"] = "firman"

	payload, _ := json.Marshal(&data)

	bodyReq := bytes.NewBuffer(payload)

	req, err := http.NewRequest("POST", "http://localhost:8081/test-with-timeout", bodyReq)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	defer res.Body.Close()

	fmt.Printf("res.Status: %v\n", res.Status)
	fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
	var resBody any

	decod := json.NewDecoder(res.Body)
	decod.Decode(&resBody)
	fmt.Printf("resBody: %v\n", resBody)

}
