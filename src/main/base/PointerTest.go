package base

import "fmt"

func pointerTest() {
	i1 := 5
	p := &i1
	var p2 *int = p
	// %p 表示指针格式，用十六进制表示
	fmt.Printf("%d, %p, %d, %p", i1, p, p, p2)

	s := "hello"
	p3 := &s
	*p3 = "good"
	fmt.Printf("\nHere is the pointer p: %p\n", p3)     // prints address
	fmt.Printf("Here is the stringLearn *p: %s\n", *p3) // prints stringLearn
	fmt.Printf("Here is the stringLearn s: %s\n", s)    // prints same stringLearn
}
