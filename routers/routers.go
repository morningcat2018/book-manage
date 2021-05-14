package routers

import (
	"book-manage/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	//r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	//r.LoadHTMLGlob("templates/*")

	// 最基本的用法
	router.GET("/", controller.Home)
	router.GET("/home", controller.Home)
	router.GET("/hello", controller.Hello)

	// v1
	bookGroup := router.Group("book")
	{
		// 最基本的用法
		bookGroup.POST("/add", controller.AddBook) // http://localhost:9090/book/add
		bookGroup.POST("/update", controller.AddBook)
		bookGroup.GET("/getDetail", controller.GetBookDetail) // http://localhost:9090/book/getDetail?code=BN20001
		bookGroup.GET("/delete", controller.DeleteBook)
		bookGroup.GET("/getList", controller.GetBookList) // http://localhost:9090/book/getList
	}
	return router
}
