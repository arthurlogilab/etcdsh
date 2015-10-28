package commands

import "strings"
import "fmt"
import "github.com/kamilhark/etcdsh/pathresolver"
import "github.com/kamilhark/etcdsh/etcdclient"
import "github.com/kamilhark/etcdsh/common"

type SetCommand struct {
	PathResolver *pathresolver.PathResolver
	EtcdClient   etcdclient.EtcdClient
}

func (c *SetCommand) Supports(command string) bool {
	return strings.EqualFold(command, "set")
}

func (c *SetCommand) Handle(args []string) {
	key := c.PathResolver.Resolve(args[0])
	value := args[1]
	err := c.EtcdClient.Set(key, value)
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

func (c *SetCommand) CommandString() string {
	return "set"
}

func (o *SetCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available:true, OnlyDirs:true}
}


