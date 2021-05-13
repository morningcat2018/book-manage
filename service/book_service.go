package service

import (
	"book-manage/dao/mysql_impl"
	"book-manage/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

var bookDao mysql_impl.MysqlImpl

func BookAdd(context *gin.Context) {
	b, _ := context.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m entity.Book
	// 反序列化
	_ = json.Unmarshal(b, &m)
	bookDao.SaveBook(&m)
	context.JSON(http.StatusOK, "add or update success")
}

func BookGetDetail(context *gin.Context) {
	code := context.Query("code")
	book := bookDao.QueryById(code)
	context.JSON(http.StatusOK, book)
}

func BookDeleteBookByCode(context *gin.Context) {
	code := context.Query("code")
	bookDao.DeleteBook(code)
	context.JSON(http.StatusOK, "delete success")
}

func BookGetList(context *gin.Context) {
	bookList := bookDao.QueryList()
	context.JSON(http.StatusOK, bookList)
}
