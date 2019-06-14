package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoginEndpoint(ctx *gin.Context) {
	fmt.Print("This is Service Login Fun ")
	ctx.JSON(200, gin.H{
		"message": "This is Gin restful",
	})
}

func LoginFind(ctx *gin.Context) {
	//id := ctx.ParamInt("id")
	id := 1
	ctx.JSON(200, map[string]interface{}{
		"status":  "自定义结构体",
		"message": fmt.Sprintf("查询ID: %d 的数据", id),
	})
}
