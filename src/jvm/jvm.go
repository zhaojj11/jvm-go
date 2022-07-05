package jvm

import (
	"fmt"
	"jvm-go/src/cmd"
)

func StartJVM(cmd *cmd.Cmd) {
	fmt.Printf("classpath: %s class:%s args:%v\n", cmd.CpOption, cmd.Class, cmd.Args)
}
