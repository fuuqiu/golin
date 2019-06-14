package main

import (
	"api"
	_ "github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//项目启动入口
func main() {
	router := api.InitRouter()
	_ = router.Run(":8083")
}
