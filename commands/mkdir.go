package commands

import "strings"
import "fmt"
import "github.com/kamilhark/etcdsh/pathresolver"
import (
	"github.com/kamilhark/etcdsh/common"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

type MkDirCommand struct {
	PathResolver *pathresolver.PathResolver
	KeysApi      client.KeysAPI
}

func (c *MkDirCommand) Supports(command string) bool {
	return strings.EqualFold(command, "mkdir")
}

func (c *MkDirCommand) Handle(args []string) {
	key := c.PathResolver.Resolve(args[0])
	_, err := c.KeysApi.Set(context.Background(), key, "", &client.SetOptions{Dir:true})
	if err != nil {
		fmt.Println(err)
	}
}

func (c *MkDirCommand) Verify(args []string) error {
	if len(args) != 1 {
		return common.NewStringError("wrong number of arguments, mkdir command requires one argument")
	}
	return nil
}

func (c *MkDirCommand) CommandString() string {
	return "mkdir"
}

func (o *MkDirCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available:false}
}


