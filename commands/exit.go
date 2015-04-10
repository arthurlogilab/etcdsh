package commands

import "os"
import "strings"

type ExitCommand struct {
}

func NewExitCommand() *ExitCommand {
	return new(ExitCommand)
}

func (exitCommand *ExitCommand) Supports(command string) bool {
	return strings.EqualFold(command, "exit")
}

func (exitCommand *ExitCommand) Handle(args []string) {
	os.Exit(0)
}

func (exitCommand *ExitCommand) Verify(args []string) error {
	return nil
}
