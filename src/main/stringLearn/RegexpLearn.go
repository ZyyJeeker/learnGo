package stringLearn

import (
	"regexp"
	"strings"
)

func RegexTest() {
	sourceStr := "我的电话号码是15789567899，请Jack兄核对"
	pat1 := "\\d+"
	pat2 := "\\w+"

	ok1, _ := regexp.MatchString(pat1, sourceStr)
	println(ok1)
	ok2, _ := regexp.MatchString(pat2, sourceStr)
	println(ok2)

	// 替换为#号
	reg, _ := regexp.Compile(pat1)
	res := reg.ReplaceAllString(sourceStr, "####")
	println(res)

	// 替换为小写
	com, _ := regexp.Compile(pat2)
	stringFuncRes := com.ReplaceAllStringFunc(sourceStr, func(s string) string {
		return strings.ToLower(s)
	})
	println(stringFuncRes)
}
