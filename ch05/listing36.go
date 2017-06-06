package main

import "fmt"

//notifier是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

type user3 struct {
	name string
	email string
}

//notify是使用指针接收者实现的方法
func (u *user3) notify()  {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	user :=user3{"Bill", "bill@email.com"}
	//sendNotification(user)//编译错误，因为方法是使用指针接收者声明 user3 does not implement notifier (notify method has pointer receiver)
//	如果使用指针接收者来实现一个接口，那么只有指向那个类型的指针才能实现对应的接口
//	如果使用值接收者来实现一个接口，那么那个类型对值和指针都实现对应的接口

	sendNotification(&user)
}

//sendNotification接受一个实现了notifier接口的值
func sendNotification(n notifier)  {
	n.notify()
}
