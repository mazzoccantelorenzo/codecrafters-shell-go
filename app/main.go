package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// BUILTIN_COMMANDS is the list of all existing and valid BUILTIN_COMMANDS
var BUILTIN_COMMANDS = []string{EXIT_COMMAND, ECHO_COMMAND, TYPE_COMMAND}

// List of BUILTIN_COMMANDS

var EXIT_COMMAND = "exit"
var ECHO_COMMAND = "echo"
var TYPE_COMMAND = "type"

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

func main() {

	for {
		PrintDollar()

		// input is the user's input
		input := textInput()

		executeProgram(input)
	}
}
