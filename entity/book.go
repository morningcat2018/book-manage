package entity

import (
	"fmt"
	"strconv"
)

/**
CREATE TABLE `book` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `book_code` VARCHAR(50) DEFAULT '',
 	`book_name` VARCHAR(100) DEFAULT '',
 	`author` VARCHAR(50) DEFAULT '',
    `publish_year` INT(11) DEFAULT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
*/

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
