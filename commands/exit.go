package commands

import "os"
import (
	"strings"
	"github.com/peterh/liner"
)

type ExitCommand struct {
	State *liner.State
}

func NewExitCommand(state *liner.State) *ExitCommand {
	exitCommand := new(ExitCommand)
	exitCommand.State = state;
	return exitCommand
}

func (exitCommand *ExitCommand) Supports(command string) bool {
	return strings.EqualFold(command, "exit")
}

func (exitCommand *ExitCommand) Handle(args []string) {
	exitCommand.State.Close()
	os.Exit(0)
}

func (exitCommand *ExitCommand) Verify(args []string) error {
	return nil
}

func (exitCommand *ExitCommand) CommandString() string {
	return "exit"
}
