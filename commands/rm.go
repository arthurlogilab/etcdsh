package commands

import "strings"

import (
	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/engine"
)

type RmCommand struct {
	Engine engine.Engine
}

func (c *RmCommand) Supports(command string) bool {
	return strings.EqualFold(command, "rm")
}

func (c *RmCommand) Handle(args []string) {
	for i := 0; i < len(args); i++ {
		c.Engine.Rm(args[i], false)
	}
}

func (c *RmCommand) Verify(args []string) error {
	if len(args) < 1 {
		return common.NewStringError("wrong number of arguments, rm command requires at least one argument")
	}
	return nil
}

func (c *RmCommand) CommandString() string {
	return "rm"
}

func (o *RmCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: false}
}
