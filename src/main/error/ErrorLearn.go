package error

import (
	"fmt"
	"log"
	"os"
)

type Any interface {
}

// ErrorTest  panic 抛异常逻辑，异常抛出到当前线程，停止当前线程，defer仍然会被执行
func ErrorTest() Any {
	var user = os.Getenv("USER")
	fmt.Println(user)
	defer func() {
		fmt.Println("发生一个错误!")
	}()
	if user == "" {
		panic("Unknown user: no value for $USER") // panic 后的代码都不会被执行，直到遇到recover，会从recover处开始继续执行
	}
	return "good"
}

func Protect(g func() Any) {
	defer func() {
		log.Println("Done")
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("Start")
	fmt.Println(g())
}
