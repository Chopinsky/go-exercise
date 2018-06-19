package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"./commands"
)

var ps1 = "$ "

func main() {
	// shell variables
	var command string
	var args []string
	var cleanInput string

	reader := bufio.NewReader(os.Stdin)

	root, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for {
		fmt.Printf(ps1)

		input, _ := reader.ReadString('\n')
		cleanInput = strings.TrimSpace(input)

		if strings.Compare(cleanInput, "exit") == 0 || strings.Compare(cleanInput, "opt") == 0 {
			break
		} else if strings.Compare(cleanInput, "") == 0 {
			continue
		} else {
			command, args = parseInput(cleanInput)
			executeCommand(command, args, &root)
		}
	}

	fmt.Printf("\nYou're leaving shell...\n")
	os.Exit(0)
}

func parseInput(args string) (string, []string) {
	if args == "" {
		return "", []string{}
	}

	argsArr := strings.Split(strings.ToLower(args), " ")
	return argsArr[0], argsArr[1:]
}

func executeCommand(command string, args []string, rootDir *string) {
	switch command {
	case "ls":
		if len(args) == 0 {
			// default to current directory
			cmd.LS(*rootDir)
		} else if len(args[0]) > 0 {
			// going to the directory to tour
			cmd.LS(args[0])
		} else {
			fmt.Printf("Invalid ls command arguments!")
		}

	case "pwd":
		cmd.PWD(*rootDir)

	default:
		fmt.Println("Invalid command!")

	}
}
