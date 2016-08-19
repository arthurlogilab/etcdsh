package commands

import (
	"fmt"
	"strings"

	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/engine"
)

type LsCommand struct {
	Engine engine.Engine
}

func (c *LsCommand) Supports(command string) bool {
	return strings.EqualFold(command, "ls")
}

func (c *LsCommand) Handle(args []string) {
	var lsArg = ""
	if len(args) == 1 {
		lsArg = args[0]
	}
	node := c.Engine.Get(lsArg, true)

	if node != nil {
		for _, n := range node.Nodes {
			fmt.Println(n.Key)
		}
	}
}

func (c *LsCommand) Verify(args []string) error {
	if len(args) > 2 {
		return common.NewStringError("to many arguments")
	}
	return nil
}

func (c *LsCommand) CommandString() string {
	return "ls"
}

func (o *LsCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: true}
}
