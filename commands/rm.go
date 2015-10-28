package commands

import "strings"
import "fmt"
import "github.com/kamilhark/etcdsh/pathresolver"
import "github.com/kamilhark/etcdsh/etcdclient"
import "github.com/kamilhark/etcdsh/common"

type RmCommand struct {
	PathResolver *pathresolver.PathResolver
	EtcdClient   etcdclient.EtcdClient
}

func (c *RmCommand) Supports(command string) bool {
	return strings.EqualFold(command, "rm")
}

func (c *RmCommand) Handle(args []string) {
	for i := 0; i < len(args); i++ {
		key := c.PathResolver.Resolve(args[i])
		err := c.EtcdClient.Delete(key)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *RmCommand) Verify(args []string) error {
	if len(args) < 1 {
		return common.NewStringError("wrong number of arguments, rm command requires at least one argument")
	}
	return nil
}

func (c *RmCommand) CommandString() string {
	return "rm"
}

func (o *RmCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available:true, OnlyDirs:false}
}

