package goroutines

import (
	"fmt"
	"time"
)

func BufChannelTest() {
	syncChan := make(chan int)
	go Sum(syncChan, 5)
	sum := <-syncChan
	fmt.Println(sum) // 阻塞的，直到Sum计算完才打印
}

func BufChannelTest1() {
	bufSize := 100
	asyncChan := make(chan int, bufSize)
	// bufSize == 0 -> synchronous, unbuffered (阻塞）
	// bufSize > 0 -> asynchronous, buffered（非阻塞）
	// 缓存大小为0表示管道的同步通信，是阻塞的
	// 缓存管道大于0表示管道是异步通信，是非阻塞的
	fmt.Println(cap(asyncChan))

	go Sum(asyncChan, 5) // 发送在满载之前不会阻塞
	sum := <-asyncChan   // 接收在空载之前不会阻塞
	fmt.Println(sum)
}

func Sum(ch chan int, num int) {
	var sum int
	for i := 0; i < num; i++ {
		sum = i + sum
		time.Sleep(1e9)
	}
	ch <- sum
}
