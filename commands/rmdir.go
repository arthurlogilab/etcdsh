package commands

import "strings"
import "fmt"
import "github.com/kamilhark/etcdsh/pathresolver"
import (
	"github.com/kamilhark/etcdsh/common"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

type RmDirCommand struct {
	PathResolver *pathresolver.PathResolver
	KeysApi      client.KeysAPI
}

func (c *RmDirCommand) Supports(command string) bool {
	return strings.EqualFold(command, "rmdir")
}

func (c *RmDirCommand) Handle(args []string) {
	for i := 0; i < len(args); i++ {
		key := c.PathResolver.Resolve(args[i])
		_, err := c.KeysApi.Delete(context.Background(), key, &client.DeleteOptions{Recursive: true, Dir: true})
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *RmDirCommand) Verify(args []string) error {
	if len(args) < 1 {
		return common.NewStringError("wrong number of arguments, rm command requires at least one argument")
	}
	return nil
}

func (c *RmDirCommand) CommandString() string {
	return "rmdir"
}

func (o *RmDirCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available:true, OnlyDirs:true}
}

