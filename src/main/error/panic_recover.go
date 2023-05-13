package error

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func RecoverTest() {
	var examples = []string{
		"one two",
		"six nxine",
		"four one",
	}
	for i := range examples {
		words := examples[i]
		numbers, err := Parse(words)
		if err != nil {
			log.Printf("解析错误:%s", words)
		}
		fmt.Println(numbers)
	}
	s := func(s Any) Any {
		panic(errors.New("make error on purpose"))
	}
	wrapFunc := RecoverWrap(s)
	wrapFunc("")
	fmt.Println("Done! ### Done!")
}

var words2NumberMap = make(map[string]int)

func init() {
	words2NumberMap["one"] = 1
	words2NumberMap["two"] = 2
	words2NumberMap["three"] = 3
	words2NumberMap["four"] = 4
	words2NumberMap["five"] = 5
	words2NumberMap["six"] = 6
	words2NumberMap["seven"] = 7
	words2NumberMap["eight"] = 8
	words2NumberMap["nine"] = 9
}

type ParseError struct {
	Index int
	Word  string
}

func (p ParseError) Error() string {
	return fmt.Sprintf("Error occured at %d index, and the word occured error is %q", p.Index, p.Word)
}

func Parse(words string) ([]int, error) {
	defer func() {
		if err := recover(); err != nil {
			if err2, ok := err.(error); ok {
				log.Printf("run time panic: %v", err2.Error())
			}
		}
	}()
	fields := strings.Fields(words)
	return fields2Numbers(fields), nil
}

func fields2Numbers(fields []string) []int {
	if len(fields) == 0 {
		panic(ParseError{
			Index: -1,
			Word:  "",
		})
	}
	var res []int
	for idx, field := range fields {
		i := words2NumberMap[field]
		if i > 0 && i < 10 {
			res = append(res, i)
		} else {
			panic(ParseError{idx, field})
		}
	}
	return res
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// RecoverWrap recover包装函数
func RecoverWrap(f func(any2 Any) Any) func(any2 Any) Any {
	return func(a Any) Any {
		defer func() {
			if err, ok := recover().(error); ok {
				log.Printf("run time panic: %v", err)
			}
		}()
		return f(a)
	}
}
