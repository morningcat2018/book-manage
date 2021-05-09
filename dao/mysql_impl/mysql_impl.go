package mysql_impl

import (
	"book-manage/entity"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_go_test?charset=utf8mb4&parseTime=True"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

type MysqlImpl struct{}

func (m MysqlImpl) SaveBook(book *entity.Book) {
	sqlStr := "insert into book(book_code,book_name,author, publish_year) values (?,?,?,?)"
	ret, err := db.Exec(sqlStr, book.BookCode, book.BookName, book.Author, book.PublishYear)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func (m MysqlImpl) DeleteBook(bookId string) {

}

func (m MysqlImpl) QueryById(bookCode string) *entity.Book {
	sqlStr := "select book_code,book_name,author, publish_year from book where book_code=?"
	var book entity.Book
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&book.BookCode, &book.BookName, &book.Author, &book.PublishYear)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return nil
	}
	return &book
}

func (m MysqlImpl) QueryListByName(bookName string) []entity.Book {
	sqlStr := "select book_code,book_name,author, publish_year from book where book_name=?"
	rows, err := db.Query(sqlStr, bookName)
	return list(err, rows)
}

func (m MysqlImpl) QueryList() []entity.Book {
	sqlStr := "select book_code,book_name,author, publish_year from book"
	rows, err := db.Query(sqlStr)
	return list(err, rows)
}

func list(err error, rows *sql.Rows) []entity.Book {
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	var booList = make([]entity.Book, 5, 10)
	// select count 优化
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.BookCode, &book.BookName, &book.Author, &book.PublishYear)
		booList = append(booList, book)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil
		}
	}
	return booList
}
