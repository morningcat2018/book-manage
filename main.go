package main

import (
	"bufio"
	"container/list"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() { DataPanel() }

type Book struct {
	bookName    string
	bookCode    string
	author      string
	publishYear int
}

func (book *Book) GetBookString() string {
	var str = book.bookCode + "\t" + book.bookName + "\t" + book.author + "\t" + strconv.Itoa(book.publishYear)
	return str + "\n"
}
func (book *Book) PrintBook() {
	fmt.Printf("编号：%s\t", book.bookCode)
	fmt.Printf("书名：%s\t", book.bookName)
	fmt.Printf("作者：%s\t", book.author)
	fmt.Printf("出版时间：%d\n", book.publishYear)
}

var (
	folderName = "data/"
)

func DataPanel() {
OUT:
	for {
		fmt.Print(getPanelString())
		fmt.Print("请输入指令：")
		var commond string
		fmt.Scanln(&commond)
		switch commond {
		case "1":
			inputBookInfo()
			break
		case "2":
			queryBook()
			break
		case "3":
			deleteBookByCode()
			break
		case "4":
			printAllBookList()
			break
		default:
			fmt.Println("再会")
			break OUT
		}
	}
}
func printAllBookList() {
	books := list.New()
	files, _ := ioutil.ReadDir("./" + folderName)
	for _, f := range files {
		queryBookByFile(folderName+f.Name(), "", true, books)
	}
	printBookList(books)
}
func getPanelString() string {
	var content = "\n"
	content += "\t1. 录取新书\n"
	content += "\t2. 查找书籍\n"
	content += "\t3. 删除数据\n"
	content += "\t4. 打印数据\n"
	content += "\t0. 退出\n"
	content += "\n"
	return content
}
func newBook(code, name, author string, publishYear int) *Book {
	p := Book{bookName: name, bookCode: code, author: author, publishYear: publishYear}
	return &p
}
func printBookList(books *list.List) {
	count := 0
	for e := books.Front(); e != nil; e = e.Next() {
		book := e.Value.(*Book)
		book.PrintBook()
		count++
	}
	if count == 0 {
		fmt.Println("暂无书籍")
	}
}
func inputBookInfo() {
	var code, name, author string
	var year int
	fmt.Print("请输入书籍编码：")
	fmt.Scanln(&code)
	fmt.Print("请输入书名：")
	fmt.Scanln(&name)
	fmt.Print("请输入作者：")
	fmt.Scanln(&author)
	fmt.Print("请输入出版时间：")
	fmt.Scanln(&year)
	newBook := newBook(code, name, author, year)
	writeToFile("", newBook)
}
func queryBook() {
	var code string
	fmt.Print("请输入编码：")
	fmt.Scanln(&code)
	var fileName = getHashCode(code)
	queryBookList := list.New()
	queryBookByFile(folderName+fileName, code, true, queryBookList)
	if queryBookList.Len() > 0 {
		fmt.Println("查找到的书籍为：")
		printBookList(queryBookList)
	}
}
func queryBookByFile(fileName string, code string, isQuery bool, books *list.List) {
	inputFile, inputError := os.Open(fileName)
	if inputError != nil {
		return
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if inputString != "\n" && len(inputString) > 0 {
			bookInfo := strings.Split(inputString, "\t")
			if code != "" && ((isQuery && bookInfo[0] != code) || (!isQuery && bookInfo[0] == code)) {
				continue
			}
			year := strings.Replace(bookInfo[3], "\r", "", -1)
			year = strings.Replace(year, "\n", "", -1)
			publishYear, err := strconv.Atoi(year)
			if err != nil {
				fmt.Printf("时间数据格式不正确\n")
				continue
			}
			book := newBook(bookInfo[0], bookInfo[1], bookInfo[2], publishYear)
			books.PushBack(book)
		}
		if readerError == io.EOF {
			break
		}
	}
}
func deleteBookByCode() {
	var code string
	fmt.Print("请输入书籍编码：")
	fmt.Scanln(&code)
	var fileName = getHashCode(code)
	queryBookList := list.New()
	queryBookByFile(folderName+fileName, code, false, queryBookList)
	os.Remove(folderName + fileName)
	if queryBookList.Len() > 0 {
		writeListToFile(fileName, queryBookList)
	}
	fmt.Println("删除完成")
}
func writeToFile(hashCode string, book *Book) {
	var fileName = hashCode
	if hashCode == "" {
		fileName = getHashCode(book.bookCode)
	}
	books := list.New()
	books.PushBack(book)
	writeListToFile(fileName, books)
}
func writeListToFile(hashCode string, books *list.List) {
	if books.Len() < 1 {
		return
	}
	var fileName = hashCode
	outputFile, outputError := os.OpenFile(folderName+fileName, os.O_APPEND|os.O_WRONLY, 0666)
	if outputError != nil {
		panic("打开文件失败" + outputError.Error())
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	for e := books.Front(); e != nil; e = e.Next() {
		book := e.Value.(*Book)
		outputWriter.WriteString(book.GetBookString())
	}
	flushError := outputWriter.Flush()
	if flushError != nil {
		fmt.Println(flushError.Error())
	}
}
func getHashCode(input string) (code string) {
	hash := md5.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	result := hash.Sum(nil)
	code = fmt.Sprintf("%x", result)
	return
}
