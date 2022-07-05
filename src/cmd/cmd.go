package cmd

import (
	"flag"
	"fmt"
	"os"
)

// Cmd java 命令行
// java [-options] Class [Args]
// java [-options] -jar jar-file [Args]
// -version / -? / -help / -cp / -classpath / -Dproperty=value / -Xmx<size>...
type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	Args        []string
}

func ParseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}

	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] Class [Args...]\n", os.Args[0])
}
