package commands

import "fmt"
import "github.com/kamilhark/etcdsh/common"
import "github.com/kamilhark/etcdsh/pathresolver"
import (
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

type DumpCommand struct {
	PathResolver *pathresolver.PathResolver
	KeysApi      client.KeysAPI
}

func (c *DumpCommand) Supports(command string) bool {
	return command == "dump"
}

func (c *DumpCommand) Handle(args []string) {
	var lsArg = ""
	if len(args) == 1 {
		lsArg = args[0]
	}
	lsPath := c.PathResolver.Resolve(lsArg)
	resp, err := c.KeysApi.Get(context.Background(), lsPath, &client.GetOptions{
		Recursive: true,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	recurseDump(resp.Node)
}

func recurseDump(n *client.Node) {
	if n.Dir {
		fmt.Printf("%s/\n", n.Key)
		for _, node := range n.Nodes {
			recurseDump(node)
		}
	} else {
		fmt.Printf("%s#%s\n", n.Key, n.Value)
	}
}

func (c *DumpCommand) Verify(args []string) error {
	if len(args) > 2 {
		return common.NewStringError("to many arguments")
	}
	return nil
}

func (c *DumpCommand) CommandString() string {
	return "dump"
}

func (o *DumpCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: true}
}
