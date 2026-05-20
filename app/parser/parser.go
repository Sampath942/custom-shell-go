package parser

import (
	"strings"
)

func ParseInput(input string) (string, string) {
	parts := strings.Split(input, " ")
	command, args := parts[0], strings.Join(parts[1:], " ")
	return command, args
}
