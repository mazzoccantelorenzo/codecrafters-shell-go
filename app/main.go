package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Commands is the list of all existing and valid commands

var commands = []string{EXIT_COMMAND, ECHO_COMMAND}

// List of commands

var EXIT_COMMAND = "exit"
var ECHO_COMMAND = "echo"

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

	return slices.Contains(commands, input)

}
func getCommandFromInput(input string) string {
	// Command is the first slice of the entire string
	// 'echo hello world' ---> echo is command

	return strings.Split(input, " ")[0]
}

func printCommandNotFound(command string) {
	fmt.Printf("%s: command not found\n", command)
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
			textToPrint := strings.Split(input, command+" ")[1]
			fmt.Println(textToPrint)

		}

		//------- Input here is not valid -------------------

		if !commandIsValid(command) {
			printCommandNotFound(command)
		}

	}
}
