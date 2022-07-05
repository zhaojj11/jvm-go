package main

import (
	"fmt"
	"jvm-go/src/cmd"
	"jvm-go/src/jvm"
)

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Println("jvm-go version 0.0.1")
	} else if command.HelpFlag || command.CpOption == "" {
		cmd.PrintUsage()
	} else {
		jvm.StartJVM(command)
	}
}
