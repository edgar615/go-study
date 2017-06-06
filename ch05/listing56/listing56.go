package main

import "fmt"

//notifier是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

//notify是使用指针接收者实现的方法
func (u *user) notify()  {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

//通过嵌入类型，与内部类型相关的标识符会提升到外部类型上．这些被提升的标识符就像直接声明在外部类型里的标识符一样,也是外部类型的一部分
//这样外部类型就组合了内部类型包含的所有属性，并且可以添加新的字段和方法．
//外部类型也可以通过声明与内部类型标识符同名的标识符来覆盖内部标识符的字段或者方法，这就是扩展和修改已有类型的方法
type admin struct {
	user //嵌入类型
	level string
}

func main() {
	ad := admin{
		user:user{
			name:"john",
			email:"john@email.com",
		},
		level:"super",
	}
	//由于内部类型的提升，内部类型实现的接口会自动提升到外部类型，这意味着由于内部类型的实现，外部类型也同样实现了这个接口
	sendNotification(&ad)
}

func sendNotification(n notifier)  {
	n.notify()
}