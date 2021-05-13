package service

import (
	"book-manage/dao/mysql_impl"
	"book-manage/entity"
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
