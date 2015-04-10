package commands

import "os"
import "strings"

type Command interface {
	Supports(string) bool
	Handle([]string)
}

type ExitCommand struct {
}

func (exitCommand *ExitCommand) Supports(command string) bool {
	return strings.EqualFold(command, "exit")
}

func (exitCommand *ExitCommand) Handle(args []string) {
	os.Exit(0)
}
