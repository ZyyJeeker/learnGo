package main

import "fmt"

func forTest(l int) {
	for a := 0; a < l; a++ {
		fmt.Printf("第%d个数字\n", a)
	}
	var a = "Hello Tom"
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c\n", a[i])
	}

	x := 5
	for x > 0 {
		fmt.Printf("%d", x)
		x--
	}

	y := 5
	for y > 0 {
		fmt.Printf("%d", y)
		y--
	}
	// 死循环
	count := 0
	for {
		if count > 5 {
			break
		}
		println("good")
		count++
	}
}

func rangeTest() {
	str := "Go is beautiful language"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}

	str2 := "Chinese: 日本語"
	for index, r := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, r, r, r, []byte(string(r)))
	}
}

func mockFor() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
