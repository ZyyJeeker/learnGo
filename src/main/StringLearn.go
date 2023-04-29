package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func stringTest() {
	str1 := "good boy \n good girl \t good dog "
	println(str1)
	str2 := `good boy \n good girl \t good dog `
	println(str2)
	println(len(str2))
	str3 := "A good good study, day day up"
	fmt.Printf("%c\n", str3[0])

	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}

func testPrintf() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point

	str3 := "Good"
	isDigit := unicode.IsDigit(int32(str3[0]))
	isLower := unicode.IsLower(int32(str3[0]))
	isUpper := unicode.IsUpper(int32(str3[0]))
	println(isDigit, isLower, isUpper)
}

func appendTest() {
	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”
}

func stringsLearn() {
	str := "This is an example of a string"
	fmt.Printf("T/F Does the string \"%s\" have prefix %s?\n", str, "Th")
	println(strings.HasPrefix(str, "Th"))
	println(strings.HasSuffix(str, "ing"))
	println(strings.Contains(str, "example"))
	println(strings.Index(str, "example"))
	println(strings.LastIndex(str, "a"))
	println(strings.Replace(str, "is", "IS", -1))
	println(strings.Count(str, "is"))
	println(strings.Repeat("Good", 3))
	println(strings.ToTitle(str))
	println(strings.ToLower(str))
	println(strings.ToUpper(str))

	str1 := " Good, New, Boy | Girl "
	println(strings.TrimSpace(str1))
	println(strings.TrimLeft(str1, " "))
	println(strings.TrimRight(str1, " "))
	fmt.Printf("%v", strings.Split("A,B,C", ","))

	strArr := []string{"Tom", "Jerry"}
	fmt.Printf("\n%v\n", strArr)
	println(strings.Join(strArr, ","))

	var orig string = "666"
	var an int
	var newS string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
