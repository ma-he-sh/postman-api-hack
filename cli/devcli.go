package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cmdHelp := flag.NewFlagSet("help", flag.ExitOnError)
	cmdHelpSection := cmdHelp.String("section", "", "section")

	if len(os.Args) < 2 {
		fmt.Println("Require command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		cmdHelp.Parse(os.Args[2:])
		fmt.Println(" section:", *cmdHelpSection)
	default:
		fmt.Println("Command not found")
		os.Exit(1)
	}
}
