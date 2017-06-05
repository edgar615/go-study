package main

import "fmt"
//定义一个结构类型
type user2 struct {
	name string
	email string
}

//notify使用值接收者实现了一个方法,如果使用值接收者声明方法，调用时会使用这个值的一个副本来执行，notify2的例子,(即使传入的是指针，也会被解引用*bill为值)
func (u user2) notify()  {
	fmt.Printf("Sending User Email To %s<%s>\n",
	u.name,
	u.email)
}

func (u user2) notify2()  {
	u.name = "edgar"
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

//changeEmail使用指针接收者实现了一个方法
//当调用使用指针接收者声明的方法时，这个方法会共享调用方法时接收者所指向的值
//即使传入的是值，也会被先引用这个值得到引用 &bill
func (u *user2) changeEmail(email string)  {
	u.email = email
}

func main() {
	//使用user类型的值可以调用使用值接收者声明的方法
	bill := user2{"Bill", "bill@email.com"}
	bill.notify()
	//使用user类型值的指针可以调用使用值接收者声明的方法
	lisa := &user2{"Lisa", "lisa@email.com"}
	lisa.notify()

	//user类型的值可以调用指针接收者声明的方法
	bill.changeEmail("bill@gmail.com")
	bill.notify()

	//user类型值的指针可以调用指针接收者声明的方法
	lisa.changeEmail("lisa@gmail.com")
	lisa.notify()

	//notify2中虽然改变了方法内部bill的name，但是对方法外部的bill却不起作用（值传递）
	bill.notify2()
	fmt.Printf("1Sending User Email To %s<%s>\n",
		bill.name,
		bill.email)
	lisa.notify2()
	fmt.Printf("1Sending User Email To %s<%s>\n",
		lisa.name,
		lisa.email)
}
