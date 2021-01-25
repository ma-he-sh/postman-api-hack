package main

import (
	"flag"
	"fmt"
	actions "github.com/ma-he-sh/postman-api-hack/cli/actions"
	"os"
)

func main() {
	cmdHelp := flag.NewFlagSet("help", flag.ExitOnError)
	cmdCode := flag.NewFlagSet("gitignore", flag.ExitOnError)

	cmdCodeSection := cmdCode.String("code", "none", "Define code to add gitgnore")
	cmdCodeAction := cmdCode.String("action", "list", "Define the action {list|insert|append}")

	if len(os.Args) < 2 {
		fmt.Println(">> Require command :: devcli")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		cmdHelp.Parse(os.Args[2:])
		actions.ShowHelp()
	case "gitignore":
		cmdCode.Parse(os.Args[2:])
		actions.GitActions(*cmdCodeAction, *cmdCodeSection)
	default:
		fmt.Println("Command not found")
		os.Exit(1)
	}
}
