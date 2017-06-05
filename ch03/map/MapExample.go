package main

import "fmt"

//映射存储了无序的键值对，类似java的map
func main() {
	//	1.创建和初始化
	//使用make创建一个映射，键的类型是string，值的类型是int
	dict1 := make(map[string]int)
	fmt.Println(dict1)
	//使用映射字面量
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict2)
	//	创建映射时，更长用的方法是使用映射字面量，映射的初始长度会根据初始化时指定的键值对的数量来确定
	//	映射的键可以是任何值，这个值的类型可以是内置的类型，也可以是结构类型，只要这个值可以使用==运算符做比较
	//	切片，函数以及包含切片的结构类型由于具有引用语义，不能作为映射的键，使用这些类型会造成编译错误

	//切片可以作为值
	dict3 := map[int][]string{}
	fmt.Println(dict3)

	//	2.使用映射
	colors := map[string]string{}
	colors["Red"] = "#da1337"

	//	nil映射不能用于存储键值对
	var colors2 map[string]string
	fmt.Println(colors2)
	//colors2["Red"] = "#da1337" //error

	//	从映射获取值，并判断键是否存在
	value, exists := colors["Red"]
	if exists {
		fmt.Println(value)
	}

	//通过键来索引映射时，即使这个键不存在也总会返回一个值,这里返回的零值
	value2 := colors["Red"]
	if value2 != "" {
		fmt.Println(value2)
	}

//迭代
	for key, value3 := range colors {
		fmt.Println(key + ":" + value3)
	}

//	删除
	delete(colors, "Red")
	fmt.Println(colors["Red"])

//	函数之间传递映射
//	在函数之间传递映射不会制造一个这个映射的副本，如果函数对这个映射做了修改时，所有对这个映射的引用都会觉察到修改

	colors3 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	for key, value3 := range colors3 {
		fmt.Println(key + ":" + value3)
	}

	foo(colors3, "Red")

	for key, value3 := range colors3 {
		fmt.Println(key + ":" + value3)
	}

}

func foo(colors map[string]string, key string)  {
	delete(colors, key)
}
