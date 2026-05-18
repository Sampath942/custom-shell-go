package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/codecrafters-io/shell-starter-go/app/executor"
	"github.com/codecrafters-io/shell-starter-go/app/parser"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		cmd = strings.TrimRight(cmd, "\r\n")
		cmd = strings.TrimLeftFunc(cmd, unicode.IsSpace)
		command, args := parser.ParseInput(cmd)
		executor.ExecuteCommand(command, args)
	}
}
