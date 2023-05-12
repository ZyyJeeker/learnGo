package IO

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var LineNumber = flag.Bool("n", false, "line number")

var num = 0

func read(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if *LineNumber {
			fmt.Printf("%d ", num)
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
		num++
		if err == io.EOF {
			break
		}
	}
	return
}

func Cat() {
	flag.Parse()
	if flag.NArg() == 0 {
		read(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		if flag.Arg(i) == "-n" {
			continue
		}
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		read(bufio.NewReader(f))
		f.Close()
	}
}
