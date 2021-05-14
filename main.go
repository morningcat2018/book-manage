package main

import (
	"book-manage/dao"
	"book-manage/model"
	"book-manage/routers"
	"fmt"
)

func main() {
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 模型绑定
	dao.MysqlDb.AutoMigrate(&model.Book{})
	// 注册路由
	router := routers.SetupRouter()
	// 绑定端口 9090
	router.Run(":9090")
}
