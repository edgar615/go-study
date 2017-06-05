package main

import "fmt"
//定义一个结构类型
type user struct {
	name string
	email string
	ext int
	privileged bool
}

//使用其他结构类型声明字段
type admin struct {
	person user
	level string
}

//	基于已有类型声明一个新类型
type Duration int64

func main() {
//	声明一个user类型的变量，初始化为零值
	var bill user
	//{  0 false}
	fmt.Println(bill)

	//声明一个user类型的变量，初始化所有字段
	lisa := user{
		name : "lisa",
		email:"lisa@email.com",
		ext:123,
		privileged:true,
	}
	//{lisa lisa@email.com 123 true}
	fmt.Println(lisa)

//	按照结构声明的顺序初始化
	edgar := user{"edgar", "edgar@email.com", 123, true}
	fmt.Println(edgar)

	fred := admin{
		person:user{
			name : "lisa",
			email:"lisa@email.com",
			ext:123,
			privileged:true,
		},
		level:"super",
	}
	//{{lisa lisa@email.com 123 true} super}
	fmt.Println(fred)

	dur := Duration(1000)
	fmt.Println(dur)

}
