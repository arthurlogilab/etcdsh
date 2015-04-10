package commands

import "strings"
import "github.com/kamilhark/etcd-console/path"
import "github.com/kamilhark/etcd-console/etcdclient"
import "github.com/kamilhark/etcd-console/common"

type CdCommand struct {
	Path       *path.EtcdPath
	etcdClient *etcdclient.EtcdClient
}

func NewCdCommand(etcdPath *path.EtcdPath, etcdClient *etcdclient.EtcdClient) *CdCommand {
	cdCommand := new(CdCommand)
	cdCommand.Path = etcdPath
	cdCommand.etcdClient = etcdClient
	return cdCommand
}

func (cdCommand *CdCommand) Supports(command string) bool {
	return strings.EqualFold(command, "cd")
}

func (cdCommand *CdCommand) Handle(args []string) {
	if len(args) == 0 {
		cdCommand.Path.Clear()
	} else {
		cdCommand.Path.Add(args[0])
	}
}

func (cdCommand *CdCommand) Verify(args []string) error {
	if len(args) > 1 {
		return common.NewStringError("'cd' command supports only one argument")
	}

	if len(args) == 0 {
		return nil
	}

	nextPath := cdCommand.Path.String() + "/" + args[0]
	_, err := cdCommand.etcdClient.Get(nextPath)

	return err
}
