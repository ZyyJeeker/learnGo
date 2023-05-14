package error

import (
	"flag"
	"fmt"
	"os"
)

var ps = flag.Bool("ps", false, "ps command")
var ll = flag.Bool("ll", false, "ll command")

func MockMain() {
	flag.Parse()
	if *ps {
		RunPs()
	}
	if *ll {
		RunLs()
	}
}

func RunLs() {
	procAttr := getProcAttr()

	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)

	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid.Pid)
}

func RunPs() {
	procAttr := getProcAttr()

	pid, err := os.StartProcess("/bin/ps", []string{"ps", "-ef"}, procAttr)

	ProcessErrorHandle(err, pid)
}

func ProcessErrorHandle(err error, pid *os.Process) {
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid.Pid)
}

func getProcAttr() *os.ProcAttr {
	env := os.Environ()

	procAttr := &os.ProcAttr{
		Env:   env,
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	return procAttr
}
