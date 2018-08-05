package router

import (
	"github.com/gin-gonic/gin"
	"gin-video/api"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	v1 := r.Group("/v")
	{
		v1.POST("/add", api.Add)
		v1.GET("/users", api.Users)
		v1.PUT("/user/:id", api.Update)
		v1.DELETE("/user/:id", api.Destroy)
	}

	return r
}
