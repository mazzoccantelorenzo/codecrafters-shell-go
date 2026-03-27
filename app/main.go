package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

// BUILTIN_COMMANDS is the list of all existing and valid BUILTIN_COMMANDS
var BUILTIN_COMMANDS = []string{EXIT_COMMAND, ECHO_COMMAND, TYPE_COMMAND, PWD_COMMAND, CD_COMMAND}

// List of BUILTIN_COMMANDS

var EXIT_COMMAND = "exit"
var ECHO_COMMAND = "echo"
var TYPE_COMMAND = "type"
var PWD_COMMAND = "pwd"
var CD_COMMAND = "cd"
var PATH = os.Getenv("PATH")

func PrintDollar() {
	fmt.Print("$ ")
}

// textInput reads the string from the user input
func textInput() string {
	// I read the string that the user writes
	// then I remove the final \n
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\n")
	return input
}

// commandIsInBuiltinCommands checks if command is valid.
// A valid command means it is in BUILTIN_COMMANDS
func commandIsInBuiltinCommands(input string) bool {
	// Check if command is valid.
	// This means that command is in existing command list or in path
	return slices.Contains(BUILTIN_COMMANDS, input)

}

// getCommandFromInput returns the command from the textInput
func getCommandFromInput(input string) string {
	// Command is the first slice of the entire string
	// 'echo hello world' ---> echo is command
	return strings.Split(input, " ")[0]
}

// getArgumentFromInput returns the argument from the textInput
func getArgumentFromInput(input string, command string) string {
	// Argument is the last slice of the entire string
	// 'echo hello world' --> hello world is argument

	textSlices := strings.Split(input, command+" ")
	if len(textSlices) < 2 {

		return ""
	} else {
		return textSlices[1]

	}
}

func printCommandNotFound(command string) {
	fmt.Print(command, ": command not found\n")
}

func printArgumentNotFound(argument string) {
	fmt.Print(argument, ": not found\n")
}

func printArgumentIsBuiltin(argument string) {
	fmt.Print(argument, " is a shell builtin\n")
}

func printArgumentIsInPath(argument string, path string) {
	fmt.Print(argument, " is ", path, "\n")
}

// commandIsInPath returns presence of command in path and its corresponding path if exists
func commandIsInPath(command string) (bool, string) {
	// Here we try to search command in PATH
	path, _ := exec.LookPath(command)

	//If given path is empty, then it's not in PATH

	if path != "" {
		return true, path
	} else {
		return false, ""
	}

}

// executeProgram executes the program and prints in shell the result of the program
func executeProgram(input string) {
	args := strings.Fields(input)

	//Variadic Unpacking to take all the arguments
	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(output))
}

func getWorkingDirectory() string {
	workingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
	}
	return workingDirectory
}

func main() {
	// command initialization
	var command = ""

	for {
		PrintDollar()

		// input is the user's input
		input := textInput()

		// command is the command that we can retrieve from the input
		command = getCommandFromInput(input)
		argument := getArgumentFromInput(input, command)

		// Check if the command is a variable.
		// It's a variable if it starts with $
		if strings.HasPrefix(command, "$") {
			command = (os.Getenv(strings.Trim(command, "$")))
		}

		switch command {

		case EXIT_COMMAND:
			return

		case PWD_COMMAND:

			// PWD returns the working directory
			workingDirectory := getWorkingDirectory()
			fmt.Println(workingDirectory)

		case ECHO_COMMAND:

			// Command is basically the first element of the user's input -> input[0]
			// We want to get the rest of the string, excluding command that is 'echo'
			// for example we split the input and get the last element
			// ("echo ", because we want the string without space at the beginning)

			textToPrint := getArgumentFromInput(input, command)
			_, _, found := strings.Cut(textToPrint, "'")
			if found {
				fmt.Print("found:", found)

				fmt.Print(textToPrint, "\n")

			} else {
				fmt.Println(textToPrint)
			}

		case CD_COMMAND:

			//CD changes the working directory
			if argument == "~" {
				argument = os.Getenv("HOME")
			}
			err := os.Chdir(argument)

			if err != nil {
				fmt.Print("cd: ", argument, ": No such file or directory\n")
			}

		case TYPE_COMMAND:

			// If the argument is a command, then we can print 'is a shell builtin'
			// For example: echo exit --> exit is the argument and it is indeed an existing command

			if commandIsInBuiltinCommands(argument) {

				printArgumentIsBuiltin(argument)

			} else {

				// I pass the argument, as the argument is the command that i want to check
				// for example echo 'cat', cat is the argument but we use it as a command usually
				isInPath, path := commandIsInPath(argument)

				if isInPath {

					printArgumentIsInPath(argument, path)

				} else {

					printArgumentNotFound(argument)

				}
			}

		default:

			//Here I check if the command is in PATH and I execute it
			commandIsInPath, _ := commandIsInPath(command)

			if commandIsInPath {
				// I pass the input as I get all the arguments inside this function
				// In this way i can pass arguments to the command called
				executeProgram(input)
			} else {

				printCommandNotFound(command)
			}
		}

	}
}
