package commands

import "strings"

import (
	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/engine"
)

type MkDirCommand struct {
	Engine engine.Engine
}

func (c *MkDirCommand) Supports(command string) bool {
	return strings.EqualFold(command, "mkdir")
}

func (c *MkDirCommand) Handle(args []string) {
	c.Engine.MkDir(args[0])
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
	return AutoCompleteConfig{Available: false}
}
