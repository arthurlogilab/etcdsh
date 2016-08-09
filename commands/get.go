package commands

import (
	"fmt"
	"strings"

	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/engine"
)

type GetCommand struct {
	Engine engine.Engine
}

func (c *GetCommand) Supports(command string) bool {
	return strings.EqualFold(command, "get")
}

func (c *GetCommand) Handle(args []string) {
	node := c.Engine.Get(args[0], false)
	if node.Dir {
		fmt.Println("dir provided, no value")
	} else {
		fmt.Println(node.Value)
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

func (o *GetCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: false}
}
