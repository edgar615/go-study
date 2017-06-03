package main

import (
	"log"
	"os"
	//_ "github.com/goinaction/code/chapter2/sample/matchers"
	//"github.com/goinaction/code/chapter2/sample/search"
)


//init函数会再main之前执行
func init()  {
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
}
