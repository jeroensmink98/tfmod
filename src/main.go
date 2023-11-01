package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current directory:", dir)

	// Check if arguments are passed
	if len(os.Args) < 2 {
		log.Fatal("Please provide the folder name as an argument")
	}

	// join the current running folder with the argument
	path := createFolder(filepath.Join(dir, os.Args[1]))
	createFilesInFolder(path, os.Args[1])
}

func createFolder(name string) string {
	// Create a new directory with 0755 permissions
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully created folder: %s\n", name)

	}
	return name
}

func createFilesInFolder(folderPath string, moduleName string) {
	filesToCreate := map[string]string{
		"variables.tf": variables_content,
		"providers.tf": "",
		"main.tf":      "",
		"outputs.tf":   "",
		"README.md":    fmt.Sprintf("# %v module", moduleName),
	}

	for fileName, fileContent := range filesToCreate {
		filePath := filepath.Join(folderPath, fileName)
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Write content to the file
		if fileContent != "" {
			_, err := file.WriteString(fileContent)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("Successfully created file: %s\n", filePath)
	}
}
