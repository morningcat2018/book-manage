package controller

import (
	"book-manage/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddBook(context *gin.Context) {
	b, _ := context.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m model.Book
	// 反序列化
	_ = json.Unmarshal(b, &m)
	model.AddBook(&m)
	context.JSON(http.StatusOK, "add or update success")
}

func GetBookDetail(context *gin.Context) {
	code := context.Query("code")
	book, _ := model.GetBookDetail(code)
	context.JSON(http.StatusOK, book)
}

func DeleteBook(context *gin.Context) {
	code := context.Query("code")
	model.DeleteBook(code)
	context.JSON(http.StatusOK, "delete success")
}

func GetBookList(context *gin.Context) {
	bookList, _ := model.GetBookList()
	context.JSON(http.StatusOK, bookList)
}
