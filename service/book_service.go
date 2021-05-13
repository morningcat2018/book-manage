package service

import (
	"book-manage/dao/mysql_impl"
	"book-manage/entity"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var bookDao mysql_impl.MysqlImpl

func BookAdd(context *gin.Context) {
	code := context.PostForm("code")
	name := context.PostForm("name")
	author := context.PostForm("author")
	year := context.PostForm("year")
	a, _ := strconv.Atoi(year)
	newBook := entity.NewBook(code, name, author, a)
	bookDao.SaveBook(newBook)
	context.JSON(http.StatusOK, "add success")
}

func BookGetDetail(context *gin.Context) {
	code := context.Query("code")
	book := bookDao.QueryById(code)
	bs, _ := json.Marshal(book)
	context.JSON(http.StatusOK, string(bs))
}

func BookDeleteBookByCode(context *gin.Context) {
	code := context.Query("code")
	bookDao.DeleteBook(code)
	context.JSON(http.StatusOK, "delete success")
}

func BookGetList(context *gin.Context) {
	bookList := bookDao.QueryList()
	bs, _ := json.Marshal(bookList)
	context.JSON(http.StatusOK, string(bs))
}
