package actions

import (
	"fmt"
	"strings"
)

// help section
func ShowHelp() {
	fmt.Println(`
		## List available gitignore files
		gitignore --action list

		## Add a gitignore file to current dir 
		gitignore --action insert --code Go

		#3 Insert multiple gitignore codes
		gitignore --action insert --code Go,Sass

		## Append another gitignore to current file
		gitignore --action append --code Sass

		## Clear current file
		gitignore --action clear
		`)
}

// Print as a list :: code source from stackoverflow
func table(input []string, cols int) {
	rows := (len(input) + cols - 1) / cols
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			i := col*rows + row
			if i >= len(input) {
				break
			}
			padding := ""
			if i < 9 {
				padding = " "
			}
			fmt.Printf("%d.%-11s%s", i+1, input[i], padding)
		}
		fmt.Println()
	}
}

// List Action
func listAction() {
	result := RequestList()
	if len(result.Data.Codes) > 0 {
		var list []string
		for _, code := range result.Data.Codes {
			list = append(list, code.Code)
		}
		table(list, 3)
	}
}

// Insert Content
func insertAction(code string) {
	codes := strings.Split(strings.TrimSpace(code), ",")
	if len(codes) > 0 {
		result := RequestCodes(codes)
		fmt.Println(result)
	} else {
		fmt.Println(">> Please define --code []")
	}
}

// actions
func GitActions(action string, code string) {
	fmt.Println("actions", action, code)
	switch {
	case action == "insert":
		if len(code) > 0 {
			insertAction(code)
		} else {
			fmt.Println(">> Please define --code []")
		}
		break
	case action == "append":
		fmt.Println("append")
		break
	case action == "clear":
		fmt.Println("clear")
		break
	case action == "list":
		listAction()
		break
	}
}
