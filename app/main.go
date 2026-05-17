package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type CommandInfo struct {
	cmdType string
	cmdFunc func(args ...interface{}) interface{}
}

var cmdMap map[string] CommandInfo
func parseInput(input string) (string, []string) {
	parts := strings.Fields(input)
	command, args := parts[0], parts[1:]
	return command, args
}

func exitFunc(args ...interface{}) interface{} {
	os.Exit(0)
	return nil
}

func echoFunc(args ...interface{}) interface{} {
	deconstructedArgs := args[0].([]string)
	fmt.Println(strings.Join(deconstructedArgs, " "))
	return nil
}

func typeFunc(args ...interface{}) interface{} {
	deconstructedArgs := args[0].([]string)
	command := deconstructedArgs[0]
	cmdInfo, exists := cmdMap[command]
	if !exists {
		fmt.Println(command + ": command not found")
	} else {
		fmt.Println(command + " is a " + cmdInfo.cmdType)
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	cmdMap = map[string] CommandInfo {
		"exit": CommandInfo{
			cmdType: "shell builtin",
			cmdFunc: exitFunc,
		},
		"echo": CommandInfo{
			cmdType: "shell builtin",
			cmdFunc: echoFunc,
		},
		"type": CommandInfo{
			cmdType: "shell builtin",
			cmdFunc: typeFunc,
		},
	}

	for {
		fmt.Print("$ ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		cmd = strings.TrimRight(cmd, "\r\n")
		cmd = strings.TrimLeftFunc(cmd, unicode.IsSpace)
		command, args := parseInput(cmd)
		cmdInfo, exists := cmdMap[command]
		if !exists {
			fmt.Println(cmd + ": command not found")
		} else {
			cmdInfo.cmdFunc(args)
		}
	}
}
