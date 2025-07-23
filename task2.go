package main

import (
	"fmt"
	"sync"
	"time"
)

func ji() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
func ou() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func runTasks(tasks []func()) {
	var wg sync.WaitGroup
	for idx, task := range tasks {
		wg.Add(1)
		//闭包
		go func(i int, t func()) {
			defer wg.Done()
			start := time.Now()
			t()
			elapsed := time.Since(start)
			fmt.Printf("任务%d 执行时间: %v\n", i+1, elapsed)
		}(idx, task)
	}
	wg.Wait()
}

func main() {
	tasks := []func(){ji, ou}
	runTasks(tasks)
}
