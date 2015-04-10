package commands

import "strings"
import "github.com/kamilhark/etcd-console/path"

type CdCommand struct {
	Path *path.EtcdPath
}

func NewCdCommand(etcdPath *path.EtcdPath) *CdCommand {
	cdCommand := new(CdCommand)
	cdCommand.Path = etcdPath
	return cdCommand
}

func (cdCommand *CdCommand) Supports(command string) bool {
	return strings.EqualFold(command, "cd")
}

func (cdCommand *CdCommand) Handle(args []string) {
	cdCommand.Path.Add(args[0])
}

func (cdCommand *CdCommand) Verify(args []string) error {
	if len(args) > 1 {
		return NewStringError("you can specify only one argument")
	}

	return nil
}
