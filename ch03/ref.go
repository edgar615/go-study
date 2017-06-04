package main

import "fmt"

func main() {
	i := 1
	var j int
	j = i
	j = 2

	fmt.Printf("i:%d, j:%d \n", i,j)

	k := new(int)
	var t *int
	*k = 1
	t = k

	*t = 2

	fmt.Printf("k:%d, t:%d \n", k,t)
	fmt.Printf("&k:%d, &t:%d \n", &k,&t)
	fmt.Printf("*k:%d, *t:%d \n", *k,*t)
}
