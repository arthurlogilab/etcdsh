package commands

import (
	"fmt"
	"strings"

	"github.com/coreos/etcd/client"
	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/engine"
)

type DumpCommand struct {
	Engine engine.Engine
}

func (c *DumpCommand) Supports(command string) bool {
	return strings.EqualFold(command, "dump")
}

func (c *DumpCommand) Handle(args []string) {
	var dumpArg = ""
	if len(args) == 1 {
		dumpArg = args[0]
	}
	node := c.Engine.Get(dumpArg, true)
	if node != nil {
		recurseDump(node)
	}
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
