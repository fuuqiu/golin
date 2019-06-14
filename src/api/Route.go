package api

import (
	. "api/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		login := v1.Group("/login")
		{
			login.POST("/", LoginEndpoint)
			login.GET("/:id", LoginFind)
		}
		room := v1.Group("/rooms")
		{
			room.GET("/:id", RoomFind)
		}
	}
	return router
}
