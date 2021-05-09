package dao

import "book-manage/entity"

type BookDao interface {
	/**
	新增 加 编辑
	*/
	SaveBook(book *entity.Book)

	/**
	删除
	*/
	DeleteBook(bookId string)

	QueryById(bookId string) *entity.Book

	QueryListByName(bookName string) []entity.Book

	QueryList() []entity.Book
}
