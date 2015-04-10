package commands

import "os"
import "strings"
import "github.com/kamilhark/etcd-console/common"

type LsCommand struct {
}

func NewLsCommand() *LsCommand {
	return new(LsCommand)
}

func (c *LsCommand) Supports(command string) bool {
	return strings.EqualFold(command, "ls")
}

func (c *LsCommand) Handle(args []string) {

}

func (c *LsCommand) Verify(args []string) error {
	if len(args) > 1 {
		return common.NewStringError("to many arguments")
	}
	return nil
}
