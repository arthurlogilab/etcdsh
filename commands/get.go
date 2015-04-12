package commands

import "strings"
import "fmt"
import "github.com/kamilhark/etcd-console/path"
import "github.com/kamilhark/etcd-console/etcdclient"
import "github.com/kamilhark/etcd-console/common"

type GetCommand struct {
	Path       *path.EtcdPath
	etcdClient *etcdclient.EtcdClient
}

func NewGetCommand(etcdPath *path.EtcdPath, etcdClient *etcdclient.EtcdClient) *GetCommand {
	cdCommand := new(GetCommand)
	cdCommand.Path = etcdPath
	cdCommand.etcdClient = etcdClient
	return cdCommand
}

func (c *GetCommand) Supports(command string) bool {
	return strings.EqualFold(command, "get")
}

func (c *GetCommand) Handle(args []string) {
	key := c.Path.String() + "/" + args[0]
	response, err := c.etcdClient.Get(key)
	if err != nil {
		fmt.Println(err)
	} else {
		if response.Node.Dir {
			fmt.Println("dir provided, no value")
		} else {
			fmt.Println(response.Node.Value)
		}

	}
}

func (c *GetCommand) Verify(args []string) error {
	if len(args) != 1 {
		return common.NewStringError("wrong number of arguments, get command requires one argument")
	}
	return nil
}
