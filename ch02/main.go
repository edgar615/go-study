package main

import (
	_ "go_study/ch02/matchers"
	"go_study/ch02/search"
	"log"
	"os"
)

//init函数会在main启动之前执行
func init() {
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
