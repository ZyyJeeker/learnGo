package goroutines

import (
	"fmt"
	. "learnGo/structTest"
	"time"
)

type any interface {
}

// ChannelTest 会发生deadlock，同一个线程不能对一个通道既发送又接受
func ChannelTest() {
	// 创建通道，通道的数据类型可以是int，string，any，func
	var ch1 chan string
	ch1 = make(chan string, 1) // 如果缓存容量不设置或者设置为0，会阻塞发送并发送死锁，缓存容量大于零就不会

	ch2 := make(chan any, 1)

	ch3 := make(chan func(), 1)

	// 通道发送数据
	s := "Hello"
	ch1 <- s

	p := &Person{
		FirstName: "张",
		LastName:  "宇",
	}
	ch2 <- p

	ch3 <- func() {
		fmt.Println("hello world")
	}

	// 从通道中取出数据
	var sf string = <-ch1
	fmt.Println(sf)

	var p2 = <-ch2
	fmt.Println(p2)

	var f func() = <-ch3
	f()
}

func ChannelTest1() {
	ch := make(chan string)
	go sendData(ch, []string{"Beijing", "Shanghai", "Guangzhou", "Shenzhen"})
	go getData(ch)
	time.Sleep(1e9)
}

func ChannelTest2() {
	out := make(chan int)
	go f1(out) // 提前准备好接受的goroutine就不会报错
	out <- 2
	time.Sleep(1e9)

	go func() {
		out <- 3
	}()
	f1(out)
}

func f1(in chan int) {
	fmt.Println(<-in)
}

func sendData(ch chan string, data []string) {
	for _, str := range data {
		ch <- str
	}
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
