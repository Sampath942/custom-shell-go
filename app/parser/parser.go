package parser

import "strings"

func ParseInput(input string) (string, []string) {
	parts := strings.Fields(input)
	command, args := parts[0], parts[1:]
	return command, args
}
