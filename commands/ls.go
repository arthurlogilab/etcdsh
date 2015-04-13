package commands

import "fmt"
import "github.com/kamilhark/etcd-console/common"
import "github.com/kamilhark/etcd-console/pathresolver"
import "github.com/kamilhark/etcd-console/etcdclient"

type LsCommand struct {
	pathResolver *pathresolver.PathResolver
	etcdClient   *etcdclient.EtcdClient
}

func NewLsCommand(pathResolver *pathresolver.PathResolver, etcdClient *etcdclient.EtcdClient) *LsCommand {
	lsCommand := new(LsCommand)
	lsCommand.pathResolver = pathResolver
	lsCommand.etcdClient = etcdClient
	return lsCommand
}

func (c *LsCommand) Supports(command string) bool {
	return command == "ls"
}

func (c *LsCommand) Handle(args []string) {
	currentPath := c.pathResolver.CurrentPath()
	resp, err := c.etcdClient.Get(currentPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, node := range resp.Node.Nodes {
		fmt.Println(node.Key)
	}
}

func (c *LsCommand) Verify(args []string) error {
	if len(args) > 1 {
		return common.NewStringError("to many arguments")
	}
	return nil
}
