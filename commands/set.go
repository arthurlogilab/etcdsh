package commands

import (
	"strings"

	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/engine"
)

type SetCommand struct {
	Engine engine.Engine
}

func (c *SetCommand) Supports(command string) bool {
	return strings.EqualFold(command, "set")
}

func (c *SetCommand) Handle(args []string) {
	c.Engine.Set(args[0], args[1])
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
	return AutoCompleteConfig{Available: true, OnlyDirs: true}
}
