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

func (u *admin) notify()  {
	fmt.Printf("Sending admin email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	ad := admin{
		user:user{
			name:"john",
			email:"john@email.com",
		},
		level:"super",
	}
	//给admin发送一个通知，接口的嵌入的内部类型实现并没有提升到外部类型
	//如果外部类型实现了notify方法，内部类型的实现就不会提升
	sendNotification(&ad)

	ad.user.notify()
	ad.notify()
}

func sendNotification(n notifier)  {
	n.notify()
}