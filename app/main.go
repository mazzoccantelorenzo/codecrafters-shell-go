package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func PrintDollar() {
	fmt.Print("$ ")
}

func main() {
	command := ""

	var EXIT_COMMAND = "exit"

	for command != EXIT_COMMAND {
		PrintDollar()

		//Commands is the list of all existing and valid commands
		commands := []string{EXIT_COMMAND}

		// Captures the user's command in the "command" variable.
		command, _ = bufio.NewReader(os.Stdin).ReadString('\n')

		//Remove the final \n to print in the same line after
		command = strings.Trim(command, "\n")

		//Command is valid only if command list contains it
		isValid := slices.Contains(commands, command)

		if !isValid {
			fmt.Printf("%s: command not found\n", command)
		}

	}
}
