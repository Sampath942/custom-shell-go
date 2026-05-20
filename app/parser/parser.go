package parser

import (
	"strings"

	"github.com/google/shlex"
)

func ParseInput(input string) (string, string) {
	parts, _ := shlex.Split(input)
	command, args := parts[0], strings.Join(parts[1:], " ")
	return command, args
}
