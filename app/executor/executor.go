package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CommandInfo struct {
	cmdType string
	cmdFunc func(args string) interface{}
}

var CmdMap map[string]CommandInfo

func exitFunc(args string) interface{} {
	os.Exit(0)
	return nil
}

func echoFunc(args string) interface{} {
	args = strings.Trim(args, " ")
	fmt.Println(args)
	return nil
}

// echo 'world     script' 'example''hello' test''shell
// echo 'script     example' 'world''test' hello''shell

func typeFunc(args string) interface{} {
	deconstructedArgs := strings.Fields(args)
	if len(deconstructedArgs) == 0 {
		os.Exit(1)
	}
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

func pwdFunc(args string) interface{} {
	dir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(dir)
	return nil
}

func cdFunc(args string) interface{} {
	deconstructedArgs := strings.Fields(args)
	if len(deconstructedArgs) == 0 {
		return nil
	}
	dir := deconstructedArgs[0]
	if dir == "~" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Chdir(homedir)
		return nil
	}
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("cd: " + dir + ": No such file or directory")
	}
	return nil
}

func ExecuteCommand(command string, args string) any {
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
		"cd": CommandInfo{
			cmdType: "a shell builtin",
			cmdFunc: cdFunc,
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
			fmt.Println("args: ", args)
			execCmd := exec.Command(command, args)
			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr
			execCmd.Run()
		}
	}
	return nil
}
