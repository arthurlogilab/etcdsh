package commands

import (
	"strings"

	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/engine"
)

type CdCommand struct {
	Engine engine.Engine
}

func (c *CdCommand) Supports(command string) bool {
	return strings.EqualFold(command, "cd")
}

func (c *CdCommand) Handle(args []string) {
	if len(args) == 1 {
		c.Engine.PathResolver.GoTo(args[0])
	} else {
		c.Engine.PathResolver.GoTo("")
	}
}

func (c *CdCommand) Verify(args []string) error {
	if len(args) > 1 {
		return common.NewStringError("'cd' command supports only one argument")
	}

	if len(args) == 0 {
		return nil
	}

	node := c.Engine.Get(args[0], false)

	if node != nil && !node.Dir {
		return common.NewStringError("not a directory")
	}

	return nil
}

func (c *CdCommand) CommandString() string {
	return "cd"
}

func (o *CdCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: true}
}
