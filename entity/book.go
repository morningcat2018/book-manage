package entity

import (
	"fmt"
	"strconv"
)

type Book struct {
	BookName    string
	BookCode    string
	Author      string
	PublishYear int
}

func (book *Book) GetBookString() string {
	var str = book.BookCode + "\t" + book.BookName + "\t" + book.Author + "\t" + strconv.Itoa(book.PublishYear)
	return str + "\n"
}
func (book *Book) PrintBook() {
	fmt.Printf("编号：%s\t", book.BookCode)
	fmt.Printf("书名：%s\t", book.BookName)
	fmt.Printf("作者：%s\t", book.Author)
	fmt.Printf("出版时间：%d\n", book.PublishYear)
}

func NewBook(code, name, author string, publishYear int) *Book {
	p := Book{BookName: name, BookCode: code, Author: author, PublishYear: publishYear}
	return &p
}
