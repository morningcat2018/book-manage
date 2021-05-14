package model

import "book-manage/dao"

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

// 关于 tag : https://www.cnblogs.com/chnmig/p/11382390.html
type Book struct {
	Id          uint32 `json:"id" db:"id"`
	BookName    string `json:"name" db:"book_name"`
	BookCode    string `json:"code" db:"book_code"`
	Author      string `json:"author" db:"author"`
	PublishYear int    `json:"year" db:"publish_year"`
}

// 自定义表名
func (Book) TableName() string {
	return "book"
}

/*
	增删改查
*/
func AddBook(book *Book) (err error) {
	err = dao.MysqlDb.Create(&book).Error
	return
}

func GetBookList() (bookList []*Book, err error) {
	if err = dao.MysqlDb.Find(&bookList).Error; err != nil {
		return nil, err
	}
	return
}

func GetBookDetail(bookCode string) (book *Book, err error) {
	book = new(Book)
	if err = dao.MysqlDb.Debug().Where("book_code=?", bookCode).First(book).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateBook(book *Book) (err error) {
	err = dao.MysqlDb.Save(book).Error
	return
}

func DeleteBook(bookCode string) (err error) {
	err = dao.MysqlDb.Where("book_code=?", bookCode).Delete(&Book{}).Error
	return
}
