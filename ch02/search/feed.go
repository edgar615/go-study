package search

import (
	"encoding/json"
	"os"
)

//定义一个常量
//go编译器可以根据赋值运算符右边的值来推导类型,因此声明常量的时候不需要指明类型
//小写字母开头，表明只能在search包中访问，不能暴露到包外
//它是通过字母大小写来控制可见性的，如果定义的常量、变量、类型、接口、结构、函数等的名称是大写字母开头表示能被其它包访问或调用（相当于public），非大写开头就只能在包内使用（相当于private，变量或常量也可以下划线开头）
const dataFile = "data/data.json"

// Feed contains information we need to process a feed.
//用于解码数据文件的结构类型，这个类型对外暴露
//类型里面声明了三个字段，每个字段的类型都是字符串，对应于数据文件各个文档的不同字段，
// 每个字段的声明最后'引号里的部分被称为标记(tag),这个标记里描述了JSON解码的元数据
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
//RetrieveFeeds没有参数，返回两个值：第一个是一个切片，其中每项指向一个feed类型的值，第二个返回值是一个error类型的值，用来表示函数是否调用成功
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.
	//打开一个文件,返回两个值：第一个是指针，指向File类型的值，第二个是error类型
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	//关键字defer会安排随后的函数在函数返回时才执行．再使用完文件后，需要主动关闭文件
	//使用defer来安排调用Close方法，可以保证这个函数一定会被执行
	//关键字defer可以缩短打开文件和关闭文件之间间隔对代码行数，有助提高代码可读性，减少错误
	defer file.Close()

	// Decode the file into a slice of pointers
	// to Feed values.
	//声明一个名字叫做feeds，值为nil的切片，这个切片包含一组指向Feed类型值的指针
	var feeds []*Feed
	//Decode方法里传入切片的指针
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}
