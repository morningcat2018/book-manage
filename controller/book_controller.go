package controller

import (
	"book-manage/service"
	"github.com/gin-gonic/gin"
)

func Book(router *gin.Engine) {
	// 最基本的用法
	router.POST("/book/add", service.BookAdd) // http://localhost:9090/book/add
	router.POST("/book/update", service.BookAdd)
	router.GET("/book/getDetail", service.BookGetDetail) // http://localhost:9090/book/getDetail?code=BN20001
	router.GET("/book/delete", service.BookDeleteBookByCode)
	router.GET("/book/getList", service.BookGetList) // http://localhost:9090/book/getList
}

/**
add
{
    "name": "golangAction",
    "code": "BN20003",
    "author": "mc3",
    "year": 2023
}

update
{
	"id": 1,
    "name": "xxx",
    "code": "xxx",
    "author": "xxx",
    "year": 2023
}
*/
