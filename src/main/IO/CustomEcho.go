package IO

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var NewLine = flag.Bool("n", false, "print newline")

var Help = flag.Bool("h", false, "print command help")

var HandleMap = map[bool]func([]string){
	true: func(args []string) {
		for _, arg := range args[1:] {
			if strings.EqualFold(arg, "-n") || strings.EqualFold(arg, "-h") {
				continue
			}
			fmt.Println(arg)
		}
	},
	false: func(args []string) {
		res := make([]string, 0)
		for _, arg := range args[1:] {
			if !strings.EqualFold(arg, "-n") && !strings.EqualFold(arg, "-h") {
				res = append(res, arg)
			}
		}
		fmt.Println(res)
	},
}

const (
	Space     = " "
	NewLian   = "\n"
	HelpWords = "-n will print new line \n-h show help"
)

// Echo 实现一个echo
func Echo() {
	flag.Parse()
	if *Help { // 如果有 -h 参数则打印帮助信息
		fmt.Println(HelpWords)
	}

	HandleMap[*NewLine](os.Args)
}
