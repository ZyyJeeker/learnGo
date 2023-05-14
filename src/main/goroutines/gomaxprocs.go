package goroutines

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

// 使用命令行指定使用的线程数目
var numCores = flag.Int("n", 2, "number of CPU cores to use")

func SetGOMAXPROCS() {
	flag.Parse()
	runtime.GOMAXPROCS(*numCores)
}

func WaitTime() {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}
