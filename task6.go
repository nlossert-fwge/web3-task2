package main

import (
	"fmt"
	"sync"
)

func main() {
	// 定义一个共享的计数器
	var counter int

	// 定义一个互斥锁
	var mu sync.Mutex

	// 定义一个 WaitGroup，用于等待所有协程完成
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程对计数器进行1000次递增操作
			for j := 0; j < 1000; j++ {
				// 在访问计数器前加锁
				mu.Lock()
				counter++
				// 访问结束后解锁
				mu.Unlock()
			}
		}()
	}

	// 等待所有协程执行完毕
	wg.Wait()

	// 输出最终的计数器值
	fmt.Printf("Final counter value: %d\n", counter)
}
