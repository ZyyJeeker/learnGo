package function

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

func funcAsArgs() {
	callBack(12, 23, add)
	s := "ABC123KK"
	index := strings.IndexFunc(s, func(r rune) bool {
		return r >= '0' && r <= '9'
	})
	println(index)
	// 调用匿名函数
	i := func(x, y int) int {
		return x * y
	}(3, 5)
	println(i)
	f()
	println(f1())
	println(f2())
	callBack(857556, 4678, sub())
}

func add(a, b int) int {
	return a + b
}

func callBack(a, b int, f func(int, int) int) {
	i := f(a, b)
	println(i)
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of structTest %T and has value %v\n", g, g)
	}
}

func f1() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func f2() (ret int) {
	defer func(x int) {
		x++
	}(ret)
	return 1
}

func sub() func(int, int) int {
	return func(x, y int) int {
		if x <= y {
			return x
		} else {
			return y
		}
	}
}

func AdderTest() {
	adder := Adder(5)
	println(adder(2))
	println(adder(3))
}

func Adder(init int) func(int) int {
	return func(delta int) int {
		init += delta
		return init
	}
}

func suffixTest() {
	suffixFunc := addSuffix(".jpg")
	println(suffixFunc("Hello"))
	println(suffixFunc("Hello.jpg"))
}

func addSuffix(suffix string) func(string) string {
	if len(suffix) <= 0 {
		return DefaultSuffixAdd
	}
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}

func DefaultSuffixAdd(s string) string {
	return s
}

func where() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}

func calculateTime(f func()) {
	start := time.Now()
	f()
	end := time.Now()
	fmt.Printf("longCalculation took this amount of time: %s\n", end.Sub(start))
}

func calculateTimeTest() {
	calculateTime(func() {
		sum := 0
		for i := 0; i < 1e6; i++ {
			sum += i
		}
	})
}
