package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
//	分配一个逻辑处理器给调度器
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for counter :=0; counter < 3; counter ++  {
			for char := 'a'; char < 'a'+26; char ++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for counter :=0; counter < 3; counter ++  {
			for char := 'A'; char < 'A'+26; char ++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
