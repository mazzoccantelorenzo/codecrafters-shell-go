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
	fmt.Print("$ ")
	var commands []string
	// Captures the user's command in the "command" variable
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	formattedCommand := string(strings.Split(command, "\n")[0])
	fmt.Print(formattedCommand)
	isValid := slices.Contains(commands, formattedCommand)

	if !isValid {
		fmt.Println(formattedCommand, ": command not found")
	}
}
