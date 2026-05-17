package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if(cmd == "exit") {
			break
		}
		fmt.Println(cmd + ": command not found")
	}
}
