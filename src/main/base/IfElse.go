package base

import (
	"fmt"
	"runtime"
	"strconv"
)

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func ifTest() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	println(prompt)

	if val, max := 10, 2; val > max {
		println(val, ">", max)
	}
	println(Abs(23))
	println(Abs(-333))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func conversion(s string) (rune, error) {
	res, err := strconv.Atoi(s)
	if err != nil {
		println(err.Error())
		return 0, err
	}
	return rune(res), nil
}
