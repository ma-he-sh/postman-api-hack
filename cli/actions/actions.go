package actions

import "fmt"

func ShowHelp(section string) {
	switch section {
	case "gitignore":

		break
	case "all":
	default:
		fmt.Println("show all help")
	}
}
