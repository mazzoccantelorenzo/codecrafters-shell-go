package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	for {
		fmt.Print("$ ")
		var commands []string
		// Captures the user's command in the "command" variable
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.Trim(command, "\n")

		isValid := slices.Contains(commands, command)
		if !isValid {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
