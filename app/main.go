package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var _ = fmt.Print

func main() {
	for {
		fmt.Print("$ ")
		var commands []string
		// Captures the user's command in the "command" variable.
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		//I remove the final \n to print in the same line after
		command = strings.Trim(command, "\n")

		isValid := slices.Contains(commands, command)
		if !isValid {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
