package stringLearn

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"learnGo/base"
	"regexp"
	"sort"
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
	student := base.Student{"张是", 23}
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

func sortTest() {
	sl1 := []int{23, 4, 5, 67, 12, 45}
	fmt.Println(sl1)
	sortIts(sl1)
	fmt.Println(sl1)
	println(sort.IntsAreSorted(sl1))
	println(sort.SearchInts(sl1, 5))
	fl1 := []float64{1.2, 0.3, 56.3, 12.4}
	fmt.Println(fl1)
	sortFloat(fl1)
	fmt.Println(fl1)
}

func sortIts(a []int) {
	sort.Ints(a)
}

func sortFloat(a []float64) {
	sort.Float64s(a)
}

func appendSlice() {
	a := []int{12, 34, 5}
	b := []int{6, 8, 45}
	res := append(a, b...)
	sort.Ints(res)
	fmt.Println(res)
	sl := append(res[:3], res[4:]...)
	fmt.Println(sl)
}

func copySlice() {
	a := []int{6, 8, 45}
	b := make([]int, 5)
	copy(b, a)
	fmt.Println(b)
}

func deleteSlice() {
	a := []int{12, 76, 34, 5, 25}
	res := append(a[:2], a[3:]...)
	fmt.Println(res)
}

func sliceExpand() {
	a := []int{12, 76, 34}
	res := append(a, make([]int, 3)...)
	fmt.Println(res)
}

func SliceInsert() {
	a := []int{12, 76, 34}
	res := append(a[:1], append([]int{2}, a[1:]...)...)
	fmt.Println(res)
	las := append(res[:1], append([]int{3, 4}, res[1:]...)...)
	fmt.Println(las)

	b := []int{6, 8, 45}
	las1 := append(las[:1], append(b, las[1:]...)...)
	fmt.Println(las1)

	fmt.Println(b[len(b)-1])
	fmt.Println(b[:len(b)-1])
}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	fileBytes, _ := ioutil.ReadFile(filename)
	b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	c := make([]byte, 0)
	for _, bytes := range b {
		c = append(c, bytes...)
	}
	for _, el := range c {
		fmt.Printf("%c,", el)
	}
	return c
}
