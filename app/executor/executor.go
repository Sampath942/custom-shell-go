package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CommandInfo struct {
	cmdType string
	cmdFunc func(args ...interface{}) interface{}
}

var CmdMap map[string]CommandInfo

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
	cmdInfo, exists := CmdMap[command]
	if exists {
		fmt.Println(command + " is " + cmdInfo.cmdType)
	} else {
		path, err := exec.LookPath(command)
		if err != nil {
			fmt.Println(command + ": not found")
		} else {
			fmt.Println(command + " is " + path)
		}
	}
	return nil
}

func pwdFunc(args ...interface{}) interface{} {
	dir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(dir)
	return nil
}

func ExecuteCommand(command string, args []string) any {
	CmdMap = map[string]CommandInfo{
		"exit": CommandInfo{
			cmdType: "a shell builtin",
			cmdFunc: exitFunc,
		},
		"echo": CommandInfo{
			cmdType: "a shell builtin",
			cmdFunc: echoFunc,
		},
		"type": CommandInfo{
			cmdType: "a shell builtin",
			cmdFunc: typeFunc,
		},
		"pwd": CommandInfo{
			cmdType: "a shell builtin",
			cmdFunc: pwdFunc,
		},
	}
	cmdInfo, exists := CmdMap[command]
	if exists {
		cmdInfo.cmdFunc(args)
	} else {
		_, err := exec.LookPath(command)
		if err != nil {
			fmt.Println(command + ": command not found")
		} else {
			execCmd := exec.Command(command, args...)
			output, err := execCmd.CombinedOutput()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Print(string(output))
		}
	}
	return nil
}
