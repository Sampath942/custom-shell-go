package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("$ ")
	cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	cmd = strings.TrimRight(cmd, "\r\n")
	fmt.Println(cmd + ": command not found")
}
