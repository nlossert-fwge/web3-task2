package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// AtomicCounter 使用原子操作实现的无锁计数器
type AtomicCounter struct {
	value int64
}

// NewAtomicCounter 创建一个新的原子计数器
func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{value: 0}
}

// Increment 原子递增操作
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Get 原子读取操作
func (c *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&c.value)
}

// RunAtomicCounterDemo 运行原子计数器演示
func RunAtomicCounterDemo() {
	fmt.Println("=== 原子操作无锁计数器演示 ===")

	// 创建计数器
	counter := NewAtomicCounter()

	// 使用 WaitGroup 等待所有协程完成
	var wg sync.WaitGroup

	// 启动10个协程
	goroutineCount := 10
	incrementPerGoroutine := 1000

	fmt.Printf("启动 %d 个协程，每个协程递增 %d 次\n", goroutineCount, incrementPerGoroutine)

	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 每个协程执行1000次递增操作
			for j := 0; j < incrementPerGoroutine; j++ {
				counter.Increment()
			}

			fmt.Printf("协程 %d 完成\n", id)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	finalValue := counter.Get()
	expectedValue := int64(goroutineCount * incrementPerGoroutine)

	fmt.Printf("\n=== 结果 ===\n")
	fmt.Printf("最终计数器值: %d\n", finalValue)
	fmt.Printf("期望值: %d\n", expectedValue)

	if finalValue == expectedValue {
		fmt.Println("✅ 测试通过！原子操作确保了数据安全")
	} else {
		fmt.Println("❌ 测试失败！存在数据竞争")
	}
}
func main() {
	RunAtomicCounterDemo()
}
