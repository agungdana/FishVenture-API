package restsvr

import "github.com/gin-gonic/gin"

func ResponsJson(c *gin.Context, res *HttpResponse) {
	c.JSON(200, res)
}
