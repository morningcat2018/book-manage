package service

import (
	"book-manage/dao"
	"book-manage/entity"
	"fmt"
)

func InputBookInfo(dao dao.BookDao) {
	var code, name, author string
	var year int
	fmt.Print("请输入书籍编码：")
	fmt.Scanln(&code)
	fmt.Print("请输入书名：")
	fmt.Scanln(&name)
	fmt.Print("请输入作者：")
	fmt.Scanln(&author)
	fmt.Print("请输入出版时间：")
	fmt.Scanln(&year)
	newBook := entity.NewBook(code, name, author, year)
	dao.SaveBook(newBook)
}

func QueryBook(dao dao.BookDao) []entity.Book {
	var name string
	fmt.Print("请输入书名：")
	fmt.Scanln(&name)
	return dao.QueryListByName(name)
}

func DeleteBookByCode(dao dao.BookDao) {
	var code string
	fmt.Print("请输入书籍编码：")
	fmt.Scanln(&code)
	book1 := dao.QueryById(code)
	if book1 == nil {
		fmt.Errorf("book not found")
		return
	}
	dao.DeleteBook(code)
}

func PrintBookDefault(dao dao.BookDao) {
	PrintBook(dao.QueryList())
}

func PrintBook(bookList []entity.Book) {
	for _, v := range bookList {
		if v.BookCode == "" {
			continue
		}
		v.PrintBook()
	}
}
