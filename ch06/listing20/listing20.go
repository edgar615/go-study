package main

import (
	"sync"
	"fmt"
	"math/rand"
)

var wg sync.WaitGroup

func main() {
	court := make(chan int)

	wg.Add(2)

	go Player("Nadal", court)
	go Player("Djokovic", court)

	//发球
	court <- 1

	wg.Wait()

}

func Player(name string, court chan int)  {
	defer wg.Done()

	for  {
		//等待球被击打过来
		ball, ok := <- court
		if !ok {
			//如果通道关闭，说明我们赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}
		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			//关闭通道，表示我们输了
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball ++

		//将球打向对手
		court <- ball
	}
}