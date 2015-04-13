package commands

import "strings"
import "fmt"
import "github.com/kamilhark/etcd-console/pathresolver"
import "github.com/kamilhark/etcd-console/etcdclient"
import "github.com/kamilhark/etcd-console/common"

type RmCommand struct {
	PathResolver *pathresolver.PathResolver
	etcdClient   *etcdclient.EtcdClient
}

func NewRmCommand(pathResolver *pathresolver.PathResolver, etcdClient *etcdclient.EtcdClient) *RmCommand {
	command := new(RmCommand)
	command.PathResolver = pathResolver
	command.etcdClient = etcdClient
	return command
}

func (c *RmCommand) Supports(command string) bool {
	return strings.EqualFold(command, "rm")
}

func (c *RmCommand) Handle(args []string) {
	key := c.PathResolver.Resolve(args[0])
	err := c.etcdClient.Delete(key)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *RmCommand) Verify(args []string) error {
	if len(args) != 1 {
		return common.NewStringError("wrong number of arguments, rn command requires one argument")
	}
	return nil
}
