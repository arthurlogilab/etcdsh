package commands

import "strings"

import (
	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/engine"
)

type RmDirCommand struct {
	Engine engine.Engine
}

func (c *RmDirCommand) Supports(command string) bool {
	return strings.EqualFold(command, "rmdir")
}

func (c *RmDirCommand) Handle(args []string) {
	for i := 0; i < len(args); i++ {
		c.Engine.Rm(args[i], true)
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
	return AutoCompleteConfig{Available: true, OnlyDirs: true}
}
