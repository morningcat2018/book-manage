package mysql_impl

import (
	"book-manage/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sql.DB 的超集
var db *sqlx.DB

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_go_test?charset=utf8mb4&parseTime=True"
	var err error
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10) // 数据库建立连接的最大数目
	db.SetMaxIdleConns(5)  // 连接池中的最大闲置连接数
}

type MysqlImpl struct{}

func (m MysqlImpl) SaveBook(book *entity.Book) int64 {
	if book.Id != 0 {
		// 更新
		sqlStr := "update book set book_code=?,book_name=?,author=?,publish_year=? where id = ?"
		ret, err := db.Exec(sqlStr, book.BookCode, book.BookName, book.Author, book.PublishYear, book.Id)
		if err != nil {
			fmt.Printf("update failed, err:%v\n", err)
			return 0
		}
		n, err := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			fmt.Printf("get RowsAffected failed, err:%v\n", err)
			return 0
		}
		fmt.Printf("update success, affected rows:%d\n", n)
		return n
	}

	sqlStr := "insert into book(book_code,book_name,author, publish_year) values (?,?,?,?)"
	ret, err := db.Exec(sqlStr, book.BookCode, book.BookName, book.Author, book.PublishYear)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return 0
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
	return theID
}

func (m MysqlImpl) DeleteBook(bookCode string) {
	sqlStr := "delete from book where book_code = ?"
	ret, err := db.Exec(sqlStr, bookCode)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func (m MysqlImpl) QueryById(bookCode string) *entity.Book {
	sqlStr := "select book_code,book_name,author,publish_year from book where book_code=?"
	var book entity.Book
	err := db.Get(&book, sqlStr, bookCode)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return nil
	}
	return &book
}

func (m MysqlImpl) QueryListByName(bookName string) []entity.Book {
	sqlStr := "select book_code,book_name,author,publish_year from book where book_name=?"
	var booList []entity.Book
	err := db.Select(&booList, sqlStr, bookName)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	return booList
}

func (m MysqlImpl) QueryList() []entity.Book {
	sqlStr := "select * from book"
	var booList []entity.Book
	err := db.Select(&booList, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	return booList
}

func list(sqlStr string, args ...interface{}) []entity.Book {
	var booList []entity.Book
	err := db.Select(&booList, sqlStr, args)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	return booList
}
