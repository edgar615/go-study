package main

import (
	"fmt"
	"go_study/ch05/listing68/counter"
)

func main() {
	counter := counter.New(10)

	// ./listing64.go:15: cannot refer to unexported name
	//                                         counters.alertCounter
	// ./listing64.go:15: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
