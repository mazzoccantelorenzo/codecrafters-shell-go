package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

// Commands is the list of all existing and valid BUILTIN_COMMANDS

var BUILTIN_COMMANDS = []string{EXIT_COMMAND, ECHO_COMMAND, TYPE_COMMAND}

// List of BUILTIN_COMMANDS

var EXIT_COMMAND = "exit"
var ECHO_COMMAND = "echo"
var TYPE_COMMAND = "type"

var PATH = os.Getenv("PATH")

func PrintDollar() {
	fmt.Print("$ ")
}

func textInput() string {

	// I read the string that the user writes
	// then I remove the final \n

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\n")
	return input
}

func commandIsValid(input string) bool {

	// Check if command is valid.
	// This means that command is in existing command list

	return slices.Contains(BUILTIN_COMMANDS, input)

}
func getCommandFromInput(input string) string {
	// Command is the first slice of the entire string
	// 'echo hello world' ---> echo is command

	return strings.Split(input, " ")[0]
}

func getArgumentFromInput(input string, command string) string {
	// Argument is the last slice of the entire string
	// 'echo hello world' --> hello world is argument
	return strings.Split(input, command+" ")[1]

}

func printCommandNotFound(command string) {
	fmt.Print(command, ": command not found\n")
}

func main() {
	// command initialization
	var command = ""

	for command != EXIT_COMMAND {
		PrintDollar()

		// input is the user's input
		input := textInput()

		// command is the command that we can retrieve from the input
		command = getCommandFromInput(input)

		//------- Input here is valid -------------------

		if command == ECHO_COMMAND {

			// Command is basically the first element of the user's input -> input[0]
			// We want to get the rest of the string, excluding command that is 'echo' for example
			// We split the input and get the last element ("echo ", because we want the string without space at the beginning)
			textToPrint := getArgumentFromInput(input, command)
			fmt.Println(textToPrint)

		}

		if command == TYPE_COMMAND {
			argument := getArgumentFromInput(input, command)
			argument = argument[:len([]byte(argument))]
			//I f the argument is a command, then we can print 'is a shell builtin'
			// For example: echo exit --> exit is the argument and it is indeed an existing command
			if slices.Contains(BUILTIN_COMMANDS, argument) {
				fmt.Print(argument, " is a shell builtin\n")
			} else if slices.Contains(BUILTIN_COMMANDS, argument) {
				// Here we try to search command in PATH
				pathDirs := strings.Split(PATH, ":")

				for _, path := range pathDirs {
					argumentPath, _ := exec.LookPath(path)

					if argumentPath != "" {
						fmt.Print(argument, " is ", argumentPath)
					}
				}

			} else {
				//If command doesn't exist, then prints not found
				fmt.Print(argument, ": not found\n")

			}

		}

		//------- Input here is not valid -------------------

		if !commandIsValid(command) {
			printCommandNotFound(command)
		}

	}
}
