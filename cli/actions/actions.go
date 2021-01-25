package actions

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var ignoreFileName = ".gitignore"

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

// get file content as bytes
func getFileContent(result []CodeData) string {
	content := ""
	for _, code := range result {
		content += code.Content
	}
	return content
}

// replace | create file
func replaceFile(fileDir string, result []CodeData) {
	fileName := fileDir + ignoreFileName
	contentData := getFileContent(result)

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(contentData)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(">> Create file:" + fileName)
	os.Exit(0)
}

// append to file
func appendFile(fileDir string, result []CodeData) {
	fileName := fileDir + ignoreFileName
	contentData := getFileContent(result)

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(contentData)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(">> Appended to file:" + fileName)
	os.Exit(0)
}

// Clear File
func clearFile(fileDir string) {
	fileName := fileDir + ignoreFileName

	f, err := os.OpenFile(fileName, os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fmt.Println(">> Cleared File:" + fileName)
	os.Exit(0)
}

// Insert Content
func insertAppendAction(action string, code string) {
	currDir := GetCurrDir()
	codes := strings.Split(strings.TrimSpace(code), ",")
	if len(codes) > 0 {
		result := RequestCodes(codes)
		if result.Type == "success" {
			if len(result.Data) > 0 {
				fileDir := currDir + "/"

				if action == "insert" {
					// replace current content
					replaceFile(fileDir, result.Data)
				} else if action == "append" {
					// append current content
					if FileExists(fileDir) {
						appendFile(fileDir, result.Data)
					} else {
						replaceFile(fileDir, result.Data)
					}
				}
			} else {
				fmt.Println(">> No result returned, request codes are case sensitive")
			}
		}
	}
	fmt.Println(">> Please define --code []")
}

// Clear Action
func clearAction() {
	currDir := GetCurrDir() + "/"

	fmt.Println(currDir)
	clearFile(currDir)
	fmt.Println(">> Cleared File")
}

// actions
func GitActions(action string, code string) {
	switch {
	case action == "insert":
		if len(code) > 0 {
			insertAppendAction(action, code)
		} else {
			fmt.Println(">> Please define --code []")
		}
		break
	case action == "append":
		if len(code) > 0 {
			insertAppendAction(action, code)
		} else {
			fmt.Println(">> Please define --code []")
		}
		break
	case action == "clear":
		clearAction()
		break
	case action == "list":
		listAction()
		break
	}
}
