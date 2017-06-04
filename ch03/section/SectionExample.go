package main

import "fmt"

//切片是一种数据结构，这种数据结构便于使用和管理数据集合．
//切片是围绕动态数组的概念构建对，可以按需自动增长和缩小．
//切片的动态增长是通过内置函数append来实现的，这个函数可以快速且高效地增长切片．
//还可以通过对切片再次切片来缩小一个切片的大小．
//因为切片的底层内存也是再连续块中分配的，所以切片还能获得索引，迭代，以及为垃圾回收优化的好处

//切片是一个很小的对象，对底层数组进行类抽象，并提供相关的操作方法．
//切片有3个字段的数据结构，这些数据结构包含go语言需要操作底层数组的元数据
//这三个字段分别是指向底层数组的指针，切片访问的元素的个数（即长度）和切片允许增长到的元素个数（即容量）
func main() {
//	１．创建切片
//	make创建切片,
//	创建一个字符串切片，长度和容量都是5个元素.如果只指定长度，那么切片的容量和长度相等
	slice := make([]string, 5)

	fmt.Println(slice)
	//创建一个长度为３个元素，容量是５个元素的整型切片
	//这时底层数组的长度是指定的容量，但是初始化后并不能访问所有的数组元素
	//切片可以访问３个元素，而底层数组拥有５个元素,剩余的２个元素可以在后期操作中合并到切片，可以通过切片访问这些元素
	//如果基于这个切片创建新的切片，新切片会和原有切片共享底层数组，也能通过后期操作来访问多余容量的元素
	//切片的容量不允许小于长度
	slice2 :=make([]int, 3, 5)

	fmt.Println(slice2)

//	切片字面量创建切片
//	长度和容量都是５
	slice3 := []string{"a","b","c","d","e"}
	fmt.Println(slice3)

	//长度和容量都是100,用空字符串初始化第１００个元素
	slice4 := []string{99:""}
	fmt.Println(slice4)

//	创建一个nil切片，可以用来描述一个不存在的切片
	var slice5 []int
	fmt.Println(slice5)

//	创建一个空切片,可以用来表示空集合
	slice6 := make([]int, 0)
	slice7 := []int{}
	fmt.Println(slice6)
	fmt.Println(slice7)

//2.赋值和切片

	slice10 := []int{10,20,30,40,50}
	slice10[1] = 25

	//对于slice[i:j]来说,新切片对长度=j-i,容量=底层数组容量-i
	slice11 := slice10[1:3]
//	两个切片共享同一个底层数组，如果一个切片修改了该底层数组的共享部分，另一个切片也可以感知到
	slice11[1] = 35
	//10,25,35,40,50
	for _, value := range slice10 {
		fmt.Println(value)
	}

//	3.增长切片
//	相对于数组而言，切片可以按需增加切片的容量
	slice12 := []int{10,20,30,40,50}
	slice13 := slice12[1:3]
	//使用原有的容量来增加一个新元素
	slice13 = append(slice13, 60)
	//20,30,60
	for _, value := range slice13 {
		fmt.Println(value)
	}
	//10,20,40,60,50
	for _, value := range slice12 {
		fmt.Println(value)
	}

	//如果切片没有足够的可用容量，append函数会创建一个新的底层数组，将被引用的现有的值复制到新数组里，再追加新值
	slice14 := []int{10,20,30,40}
	//append会智能地处理底层数组的容量增长，再切片的容量小于1000个元素时，总是会成倍地增长容量，一旦容量超过1000，容量的增长因子会设为1.25，也就是每次会增加25％的容量
	slice15 := append(slice14, 50)
	for _, value := range slice15 {
		fmt.Println(value)
	}

//	4.限制容量
	slice16 := []string{"a","b","c","d","e"}
	//新切片从底层数组引用了一个元素，容量是２
	slice17 := slice16[2:3:4]
	fmt.Println(cap(slice17))
	fmt.Println(len(slice17))

	//把长度和容量设置成一样，可以在append的时候创建一个新的底层数组，而不对原来的底层数组产生影响
	slice18 := slice16[2:3:3]
	slice18 = append(slice18, "f")

	fmt.Println("slice16")
	// a,b,c,d,e
	for _, value := range slice16 {
		fmt.Println(value)
	}
	fmt.Println("slice17")
	//c
	for _, value := range slice17 {
		fmt.Println(value)
	}
	fmt.Println("slice18")
	//c,f
	for _, value := range slice18 {
		fmt.Println(value)
	}

	slice19 := append(slice17, "f")
	fmt.Println("slice16")
	//a,b,c,f,e
	for _, value := range slice16 {
		fmt.Println(value)
	}
	fmt.Println("slice17")
	//c
	for _, value := range slice17 {
		fmt.Println(value)
	}
	fmt.Println("slice19")
	//c,f
	for _, value := range slice19 {
		fmt.Println(value)
	}

//	5.追加多个变量
	slice20 := []int{10,20,30,40}
	slice20 = append(slice20, 50, 60)
	//...运算符可以将一个切片的所有元素追加到另一个切片内
	slice20 = append(slice20, []int{70,80,90}...)
	for _, value := range slice20 {
		fmt.Println(value)
	}

//	6.迭代
//	range创建了每个元素的副本，而不是直接返回对该元素的引用
/*
value: 10 value-addr: C42000EBD8 ElemAddr: C420078100
value: 20 value-addr: C42000EBD8 ElemAddr: C420078108
value: 30 value-addr: C42000EBD8 ElemAddr: C420078110
value: 40 value-addr: C42000EBD8 ElemAddr: C420078118
value: 50 value-addr: C42000EBD8 ElemAddr: C420078120
value: 60 value-addr: C42000EBD8 ElemAddr: C420078128
value: 70 value-addr: C42000EBD8 ElemAddr: C420078130
value: 80 value-addr: C42000EBD8 ElemAddr: C420078138
value: 90 value-addr: C42000EBD8 ElemAddr: C420078140
因为迭代返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以value的地址总是相同的
*/
	for index, value := range slice20 {
		fmt.Printf("value: %d value-addr: %X ElemAddr: %X\n",
		value, &value, &slice20[index])
	}

	for index := 2; index < len(slice20); index ++ {
		fmt.Println(slice20[index])
	}

//	7.多维切片
	slice21 := [][]int{{10},{100,200}}
	slice21[0] = append(slice21[0], 20)
	fmt.Println(slice21[0][1])

//	8.函数传递
//	在函数中传递切片就是要在函数之间以值的方式传递切片，由于切片对尺寸很小，再函数间复制和传递切片对成本很低
//	复制只会复制切片本身，不会涉及底层数组
	slice22 := []int{10,20,30,40}
	slice23 := foo(slice22)
	fmt.Println("slice22")
	//10,1,30,40
	for _, value := range slice22 {
		fmt.Println(value)
	}
	fmt.Println("slice23")
	//10,1,30,40,100
	for _, value := range slice23 {
		fmt.Println(value)
	}
}

func  foo(slice []int) ([]int) {
	slice[1] = 1
	return append(slice, 100)
}

