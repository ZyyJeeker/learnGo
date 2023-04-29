package function

import (
	"fmt"
	"learnGo/base"
)

func foo(x *base.Student) {
	println()
	println(x.name)
	println(x.age)
	println(x)
}

func callFoo() {
	student := base.Student{"tom", 12}
	foo(&student)
	s := new(base.Student)
	s.name = "jack"
	s.age = 22
	println(s)
	foo(s)
}

func digNoName(x, y int) (int, int) {
	return x + y, x - y
}

func digWithName(x, y int) (j, k int) {
	j = x + y
	k = x - y
	return
}

func updateInt(p *int) {
	if *p < 0 {
		*p = -*p
	}
}

func greet(who ...string) {
	for idx, s := range who {
		fmt.Printf("第%d个人%s打招呼\n", idx, s)
	}
}

func min(arr ...int) (min int) {
	if len(arr) <= 0 {
		return 0
	}
	min = arr[0]
	defer fmt.Printf("结果为:%d\n", min)
	for _, item := range arr {
		if item < min {
			min = item
		}
	}
	fmt.Printf("结果为:%d\n", min)
	return
}
