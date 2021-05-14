package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello world!",
	})
}

func Hello(context *gin.Context) {
	name := context.Query("name")
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name,
	})
}
