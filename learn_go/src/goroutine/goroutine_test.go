package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_goroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(num int) {
			fmt.Println("num = ", num)
		}(i)
	}
	time.Sleep(time.Duration(10) * time.Second)
}

func Test_channel(t *testing.T) {
	// 默认初始化
	var chan1 chan int
	fmt.Println(chan1)
	chan1 = make(chan int)
	fmt.Println(chan1)
	close(chan1)

	// 初始化channel
	chan2 := make(chan int)
	fmt.Println("this is a test")
	// 管道操作

	// 从管道中接收值
	go rec_value(chan2)
	// 向管道中传递值
	chan2 <- 1

}

func rec_value(c chan int) {
	num := <-c
	fmt.Println("num = ", num)
}

func Test_buf(t *testing.T) {
	chan1 := make(chan int, 1)
	chan1 <- 1
	num := <-chan1
	fmt.Println("num = ", num)
}

var x int
var m sync.Mutex      // 互斥锁
var wg sync.WaitGroup // 等待组

func Test_SyncMutex(t *testing.T) {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println("x = ", x)
}
func add() {
	for i := 0; i < 5000; i++ {
		m.Lock() // 修改前加锁
		x += 1
		m.Unlock() // 修改后解锁
	}
	wg.Done()
}
