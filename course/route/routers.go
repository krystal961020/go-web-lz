package route

import (
	"github.com/gin-gonic/gin"
	"go-web-lz/route/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/aip/v1")
	{
		apiV1.GET("/getUser", v1.GetUser)
	}

	return r
}
