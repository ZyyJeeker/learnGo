package error

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var cmd = flag.String("cmd", "", "exec command")
var args = flag.String("args", "", "exec command")

func ExecCommand() {
	command := exec.Command(*cmd, *args)

	err := command.Run()

	ExecErrHandle(err)
}

func ExecErrHandle(err error) {
	if err != nil {
		fmt.Printf("Error %v exec cmd!", err) //
		os.Exit(1)
	}
}
