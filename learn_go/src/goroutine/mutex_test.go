package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	x1      int
	wg1     sync.WaitGroup
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

func Test_RWMutex(t *testing.T) {
	// 使用互斥锁，10并发写，1000并发读
	do(writeWithLock, readWithLock, 10, 1000) // x:1010 cost:16.0170156s

	// 使用读写互斥锁，10并发写，1000并发读
	do(writeWithRWLock, readWithRWLock, 10, 1000) // x:1010 cost:16.0170156s
}

// 使用互斥锁的写操作
func writeWithLock() {
	mutex.Lock() // 加互斥锁
	x += 1
	time.Sleep(10 * time.Millisecond)
	mutex.Unlock() // 解锁
	wg.Done()
}

// 使用互斥锁的写操作
func readWithLock() {
	mutex.Lock() // 加互斥锁
	x += 1
	time.Sleep(10 * time.Millisecond)
	mutex.Unlock() // 解锁
	wg.Done()

}

// 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwMutex.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutex.Unlock()
	wg.Done()
}

func readWithRWLock() {
	rwMutex.RLock()
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutex.RUnlock()
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	//  rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x, cost)

}
