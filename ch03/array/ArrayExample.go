package main

import (
	"fmt"
)

func main() {
	//声明一个包含5个元素的整形数组,所有值都为其对应的零值
	fmt.Println("var array [5]int")
	var array [5]int
	for i, a := range array {
		fmt.Printf("index:%d, value:%d \n", i, a)
	}

	//使用数组字面量声明数组
	fmt.Println("array2 := [5]int{10, 20, 30, 40, 50}")
	array2 := [5]int{10, 20, 30, 40, 50}
	for i, a := range array2 {
		fmt.Printf("index:%d, value:%d \n", i, a)
	}

	//自动计算数组的长度
	fmt.Println("array3 := [...]int{10, 20, 30, 40, 50}")
	array3 := [...]int{10, 20, 30, 40, 50}
	for i, a := range array3 {
		fmt.Printf("index:%d, value:%d \n", i, a)
	}

	//声明数组，并指定特定元素的值
	fmt.Println("array4 := [5]int{1:10,2:20}")
	array4 := [5]int{1:10,2:20}
	for i, a := range array4 {
		fmt.Printf("index:%d, value:%d \n", i, a)
	}

	//访问数组
	fmt.Println("a[i]")
	for i := 0; i < 5; i ++ {
		fmt.Printf("index:%d, value:%d \n", i, array4[i])
	}

	//使用指针
	fmt.Println("array5 := [5]*int{0:new(int), 1:new(int)}")
	array5 := [5]*int{0:new(int), 1:new(int)}
	//赋值
	*array5[0] = 10
	*array5[1] = 20
/*
	for i, a := range array5 {
		fmt.Printf("index:%d, value:%d \n", i, *a)
	}*/
	fmt.Println(*array5[0])
	fmt.Println(*array5[1])

	//数组之间赋值
	var array6 [5]string
	array7 := [5]string{"a","b","c","d","e"}
	array6 = array7

	fmt.Println("array6 = array7")
	for i, a := range array6 {
		fmt.Printf("array6:index:%d, value:%s \n", i, a)
	}
	for i, a := range array7 {
		fmt.Printf("array7:index:%d, value:%s \n", i, a)
	}

	array7[1] = "red"
	fmt.Println("array7[1] = \"red\"")
	for i, a := range array6 {
		fmt.Printf("array6:index:%d, value:%s \n", i, a)
	}
	for i, a := range array7 {
		fmt.Printf("array7:index:%d, value:%s \n", i, a)
	}

	//指针数组赋值
	var array8 [3]*string
	array9 := [3]*string{new(string),new(string),new(string)}
	*array9[0]="a"
	*array9[1]="b"
	*array9[2]="c"
	array8=array9
	fmt.Println("array8 = array9")
	for i, a := range array8 {
		fmt.Printf("array8:index:%d, value:%s \n", i, *a)
	}
	for i, a := range array9 {
		fmt.Printf("array9:index:%d, value:%s \n", i, *a)
	}

	*array8[1] = "red"
	fmt.Println("*array8[1] = \"red\"")
	for i, a := range array8 {
		fmt.Printf("array6:index:%d, value:%s \n", i, *a)
	}
	for i, a := range array9 {
		fmt.Printf("array9:index:%d, value:%s \n", i, *a)
	}

	//多维数组
	var array10 [4][2]int
	array11 := [4][2]int{{10,11},{20,21},{30,31},{40,41}}
	array10 = array11
	array12 := [4][2]int{1:{20,21}, 3:{40,41}}
	array13 := [4][2]int{1:{0:20}, 3:{1:41}}
	for i:=0; i < 4; i ++ {
		//a := array10[i]
		for j:=0; j < 2; j ++ {
			fmt.Printf("array10:index:%d,%d, value:%d \n", i, j,array10[i][j])
		}
	}
	for i:=0; i < 4; i ++ {
		//a := array10[i]
		for j:=0; j < 2; j ++ {
			fmt.Printf("array11:index:%d,%d, value:%d \n", i, j,array11[i][j])
		}
	}
	for i:=0; i < 4; i ++ {
		//a := array10[i]
		//var a = array12[i]
		for j:=0; j < 2; j ++ {
			fmt.Printf("array12:index:%d,%d, value:%d \n", i, j,array12[i][j])
		}
	}
	for i:=0; i < 4; i ++ {
		//a := array10[i]
		for j:=0; j < 2; j ++ {
			fmt.Printf("array13:index:%d,%d, value:%d \n", i, j,array13[i][j])
		}
	}

	//函数间传递数组
	array14 := [5]int{1,2,3,4,5}
	for i, a := range array14 {
		fmt.Printf("array14:index:%d, value:%d \n", i, a)
	}
	//值传递会传递一个完整的副本，会有较大的内存开销
	foo(array14)
	for i, a := range array14 {
		fmt.Printf("array14:index:%d, value:%d \n", i, a)
	}

	//传入指针，只需要复制8个字节的数据，但是值会同步修改
	foo2(&array14)
	for i, a := range array14 {
		fmt.Printf("array14:index:%d, value:%d \n", i, a)
	}

}

func  foo(array [5]int)  {
	array[1] = 10
	for i, a := range array {
		fmt.Printf("array:index:%d, value:%d \n", i, a)
	}
}

func foo2(array *[5]int)  {
	array[1] = 10
	for i, a := range array {
		fmt.Printf("array:index:%d, value:%d \n", i, a)
	}
}
