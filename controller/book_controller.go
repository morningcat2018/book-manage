package controller

import (
	"book-manage/service"
	"github.com/gin-gonic/gin"
)

func Book(router *gin.Engine) {
	// 最基本的用法
	router.POST("/book/add", service.BookAdd)
	router.POST("/book/update", service.BookAdd)
	router.GET("/book/getDetail", service.BookGetDetail) // http://localhost:9090/book/getDetail?code=BN20001
	router.GET("/book/delete", service.BookDeleteBookByCode)
	router.GET("/book/getList", service.BookGetList) // http://localhost:9090/book/getList
}
