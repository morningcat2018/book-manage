package controller

import (
	"book-manage/service"
	"github.com/gin-gonic/gin"
)

func Book(router *gin.Engine) {
	// 最基本的用法
	router.POST("/add", service.BookAdd)
	router.POST("/update", service.BookAdd)
	router.GET("/getDetail", service.BookGetDetail)
	router.GET("/delete", service.BookDeleteBookByCode)
	router.GET("/getList", service.BookGetList)
}
