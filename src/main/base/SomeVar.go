package base

import "fmt"

func baseType() {
	var a = 3
	var b = a

	h := new(Student)
	h.Name = "张三"
	h.Age = 12
	var l = h
	fmt.Println(h)

	println(a)
	println(b)

	println(&a)
	println(&b)

	println(h)
	println(l)

	println(&h)
	println(&l)
}

const NAME = "good"

// 常量用作枚举值
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const beef, two, c = "eat", 2, "veg"
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

const (
	x = iota
	y = iota
	z = iota
)

type Student struct {
	Name string
	Age  int
}
