package service

import (
	"api/dbs"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RoomFind(ctx *gin.Context) {
	id := ctx.ParamInt("id")
	room := dbs.RoomFindById(id)
	ctx.JSON(200, map[string]interface{}{
		"status":  "自定义结构体",
		"message": fmt.Sprintf("%s", room),
	})
}
