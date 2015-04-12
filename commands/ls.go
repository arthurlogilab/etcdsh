package commands

import "strings"
import "fmt"
import "github.com/kamilhark/etcd-console/common"
import "github.com/kamilhark/etcd-console/path"
import "github.com/kamilhark/etcd-console/etcdclient"

type LsCommand struct {
	path       *path.EtcdPath
	etcdClient *etcdclient.EtcdClient
}

func NewLsCommand(path *path.EtcdPath, etcdClient *etcdclient.EtcdClient) *LsCommand {
	lsCommand := new(LsCommand)
	lsCommand.path = path
	lsCommand.etcdClient = etcdClient
	return lsCommand
}

func (c *LsCommand) Supports(command string) bool {
	return strings.EqualFold(command, "ls")
}

func (c *LsCommand) Handle(args []string) {
	currentPath := c.path.String()
	ls, err := c.etcdClient.Ls(currentPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dir := range ls {
		fmt.Println(dir)
	}
}

func (c *LsCommand) Verify(args []string) error {
	if len(args) > 1 {
		return common.NewStringError("to many arguments")
	}
	return nil
}
