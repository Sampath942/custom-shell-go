package parser

import (
	"github.com/google/shlex"
)

func ParseInput(input string) (string, []string) {
	parts, _ := shlex.Split(input)
	command, args := parts[0], parts[1:]
	return command, args
}
