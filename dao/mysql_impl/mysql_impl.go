package mysql_impl

import "book-manage/entity"

type MysqlImpl struct{}

func (m MysqlImpl) SaveBook(book *entity.Book) {
	book.PrintBook()
}

func (m MysqlImpl) DeleteBook(bookId string) {

}

func (m MysqlImpl) QueryById(bookId string) *entity.Book {
	return nil
}

func (m MysqlImpl) QueryListByName(bookName string) []entity.Book {
	return nil
}

func (m MysqlImpl) QueryList() []entity.Book {
	return nil
}
