package main

import (
	"sync"
	"runtime"
	"fmt"
	"sync/atomic"
)

var (
	counter int64
	wg sync.WaitGroup
)
func main() {

	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count:= 0; count < 2; count ++ {
		atomic.AddInt64(&counter, 1)
		//当期goroutine从线程退出，并放回队列
		runtime.Gosched();

	}
}
