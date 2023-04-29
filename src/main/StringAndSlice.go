package main

import (
	"encoding/binary"
	"fmt"
	"unicode/utf8"
)

func stringSlice() {
	hello := "Hello"
	b := []byte(hello)
	for _, item := range b {
		fmt.Printf("字符为%c,Unicode码为%X", item, item)
	}
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
	c := []int32("你好")
	for _, i := range c {
		fmt.Printf("%c", i)
	}
}

func format() {
	println()
	student := Student{"张是", 23}
	fmt.Printf("张三是:%#v,他的类型为%T\n", student, student)
	fmt.Printf("张三是:%v,他的类型为%T\n", student, student)
}

func charCount(s string) int {
	// 先转换为字符切片
	sl := []rune(s)
	for _, el := range sl {
		fmt.Printf("%c, %v, byte size:%d\n", el, el, binary.Size(el))

	}
	println("等同于函数utf8.RuneCountInString(s)")
	utf8.RuneCountInString(s)
	return len(sl)
}

func stringAppend(front, after string) string {
	return string(append([]byte(front), []byte(after)...))
}

func stringAppend2(front, after string) string {
	return front + after
}

func subString(from, to int, s string) string {
	return s[from:to]
}

func updateStr(s string, from, to int, repChar byte) string {
	ar := []byte(s)
	for i := from; i <= to; i++ {
		ar[i] = repChar
	}
	return string(ar)
}

func compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1

	}
	return 0
}
