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

type admin struct {
	name string
	email string
}

func (u *admin) notify()  {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	// Create a user value and pass it to sendNotification.
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	// Create an admin value and pass it to sendNotification.
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}
func sendNotification(n notifier)  {
	n.notify()
}