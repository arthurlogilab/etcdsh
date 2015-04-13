package commands

import "strings"
import "fmt"
import "github.com/kamilhark/etcd-console/pathresolver"
import "github.com/kamilhark/etcd-console/etcdclient"
import "github.com/kamilhark/etcd-console/common"

type SetCommand struct {
	PathResolver *pathresolver.PathResolver
	etcdClient   *etcdclient.EtcdClient
}

func NewSetCommand(pathResolver *pathresolver.PathResolver, etcdClient *etcdclient.EtcdClient) *SetCommand {
	command := new(SetCommand)
	command.PathResolver = pathResolver
	command.etcdClient = etcdClient
	return command
}

func (c *SetCommand) Supports(command string) bool {
	return strings.EqualFold(command, "set")
}

func (c *SetCommand) Handle(args []string) {
	key := c.PathResolver.Resolve(args[0])
	value := args[1]
	err := c.etcdClient.Set(key, value)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *SetCommand) Verify(args []string) error {
	if len(args) != 2 {
		return common.NewStringError("wrong number of arguments, set command requires two argument")
	}
	return nil
}
