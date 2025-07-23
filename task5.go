package main

import (
	"fmt"
	"sync"
	"time"
)

type ReadonlyChannel struct {
	ch <-chan int
}

func receiveOnly(ch ReadonlyChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch.ch {
		fmt.Println("通道接收:", v)
	}
}

func sendOnly(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	go receiveOnly(ReadonlyChannel{ch}, &wg)
	go sendOnly(ch)
	wg.Wait()
	fmt.Println("所有数据已接收完毕")
}
