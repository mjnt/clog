package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Show CLI usage and exit with given exitCode
func showUsageAndExit(exitCode int) {
	fmt.Println("Usage:\n\ninit - create a new CHANGELOG.md file")
	os.Exit(exitCode)
}

// Get project name from containing folder
func getNameFromCWD() string {
	cwd, cwdError := os.Getwd()
	if cwdError != nil {
		fmt.Println(cwdError)
		os.Exit(1)
	}

	return filepath.Base(cwd)
}

// Get initial file contents given project name
func getInitFileContents(name string) string {
	fileTemplate := `# %s Change Log

All notable changes associated with a release will be documented in this file.

This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased] - YYYY-MM-DD

### Major

### Minor

### Patch

[Unreleased]: https://www.github.com/<org>/<project>/commits/master
`
	return fmt.Sprintf(fileTemplate, name)
}

// Write file contents to file path
func writeToFile(fileContents, filePath string) error {
	file, fileErr := os.Create("CHANGELOG.md")

	if fileErr != nil {
		return fileErr
	}

	_, writeErr := file.WriteString(fileContents)
	file.Close()

	return writeErr
}

func executeInit() error {
	projectName := getNameFromCWD()
	fileContents := getInitFileContents(projectName)
	fileName := "CHANGELOG.md"
	return writeToFile(fileContents, fileName)
}

// Execute a command by name
func executeCommand(name string) error {
	switch name {
	case "init":
		return executeInit()
	default:
		errorMessage := fmt.Sprintf("command: %s is not a valid command", name)
		return errors.New(errorMessage)
	}
}

// Entrypoint
func main() {

	// Parse command line arguments
	flag.Parse()

	// Print usage if no command was given
	command := flag.Arg(0)
	if command == "" {
		showUsageAndExit(0)
	}

	// Execute command
	commandErr := executeCommand(command)
	if commandErr != nil {
		log.Fatalln(commandErr)
	}
}
