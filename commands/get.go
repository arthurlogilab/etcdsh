package commands

import "strings"
import "fmt"
import "github.com/kamilhark/etcdsh/pathresolver"
import "github.com/kamilhark/etcdsh/etcdclient"
import "github.com/kamilhark/etcdsh/common"

type GetCommand struct {
	OneArgumentAutoCompleteCommand
	PathResolver *pathresolver.PathResolver
	EtcdClient   etcdclient.EtcdClient
}

func (c *GetCommand) Supports(command string) bool {
	return strings.EqualFold(command, "get")
}

func (c *GetCommand) Handle(args []string) {
	key := c.PathResolver.Resolve(args[0])
	response, err := c.EtcdClient.Get(key)
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

func (c *GetCommand) CommandString() string {
	return "get"
}

