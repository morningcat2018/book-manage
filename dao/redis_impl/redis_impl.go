package redis_impl

import (
	"book-manage/entity"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var err error
	_, err = rdb.Ping().Result()
	if err != nil {
		panic(err)
		return
	}
}

type RedisImpl struct{}

func (r RedisImpl) SaveBook(book *entity.Book) {
	str, jerr := json.Marshal(book)
	if jerr != nil {
		fmt.Printf("transaction failed, err:%v\n", jerr)
		return
	}
	err := rdb.Set(book.BookCode, str, 0).Err()
	if err != nil {
		fmt.Printf("save failed, err:%v\n", err)
		return
	}
}

func (r RedisImpl) DeleteBook(bookCode string) {
	err := rdb.Del(bookCode).Err()
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
}

func (r RedisImpl) QueryById(bookCode string) *entity.Book {
	val, err := rdb.Get(bookCode).Result()
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		panic(err)
	}
	var book entity.Book
	json.Unmarshal([]byte(val), &book)
	return &book
}

func (r RedisImpl) QueryListByName(bookName string) []entity.Book {
	var bookList []entity.Book = r.QueryList()
	for _, book := range bookList {
		if book.BookName == bookName {
			bookList = append(bookList, book)
		}
	}
	return bookList
}

func (r RedisImpl) QueryList() []entity.Book {
	keys, err := rdb.Keys("*").Result()
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		panic(err)
	}
	var bookList []entity.Book = make([]entity.Book, len(keys))
	for _, bookCode := range keys {
		val, err := rdb.Get(bookCode).Result()
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
			panic(err)
		}
		var book entity.Book
		json.Unmarshal([]byte(val), &book)
		bookList = append(bookList, book)
	}

	return bookList
}
