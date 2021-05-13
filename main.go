package main

import (
	"book-manage/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	// 注册一个默认的路由器
	router := gin.Default()
	// 绑定 home controller
	controller.Home(router)
	controller.Book(router)

	// 绑定端口 9090
	router.Run(":9090")
}
