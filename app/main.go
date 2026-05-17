package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	for {
		fmt.Print("$ ")
		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		cmd = strings.TrimLeftFunc(cmd, unicode.IsSpace)
		if cmd == "exit" {
			break
		} else if len(cmd) >= 5 && cmd[:5] == "echo " {
			remaining := cmd[5:]
			fmt.Println(remaining)
			continue
		}
		fmt.Println(cmd + ": command not found")
	}
}
