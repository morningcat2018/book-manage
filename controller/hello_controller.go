package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(router *gin.Engine) {
	// 最基本的用法
	router.GET("/", home)
	router.GET("/home", home)
	router.GET("/hello", hello) // http://localhost:9090/hello?name=morning
}

func home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}

func hello(context *gin.Context) {
	name := context.Query("name")
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name,
	})
}

func hello2(w http.ResponseWriter, r *http.Request) {

}
