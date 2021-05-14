package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	MysqlDb *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:@tcp(127.0.0.1:3306)/db_go_test?charset=utf8mb4&parseTime=True"
	MysqlDb, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return MysqlDb.DB().Ping()
}

func Close() {
	MysqlDb.Close()
}
