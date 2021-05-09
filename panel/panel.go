package panel

import (
	"book-manage/dao/mysql_impl"
	"book-manage/service"
	"fmt"
)

func DataPanel() {
	var dao mysql_impl.MysqlImpl
	for {
		fmt.Print(getPanelString())
		fmt.Print("请输入指令：")
		var commond int
		fmt.Scanln(&commond)
		if commond == 0 {
			break
		} else if commond == 1 {
			service.InputBookInfo(dao)
		} else if commond == 2 {
			bookSlice := service.QueryBook(dao)
			service.PrintBook(bookSlice)
		} else if commond == 3 {
			service.DeleteBookByCode(dao)
		} else if commond == 4 {
			service.PrintBookDefault(dao)
		}
	}
}

func getPanelString() string {
	var content = "********************\n"
	content += "\t1. 录取新书\n"
	content += "\t2. 查找书籍\n"
	content += "\t3. 删除数据\n"
	content += "\t4. 打印数据\n"
	content += "\t0. 退出\n"
	content += "********************\n"
	return content
}
