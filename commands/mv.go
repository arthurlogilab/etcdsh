package commands

import (
	"strings"

	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/engine"
)

type MvCommand struct {
	Engine engine.Engine
}

func (c *MvCommand) Supports(command string) bool {
	return strings.EqualFold(command, "mv")
}

func (c *MvCommand) Handle(args []string) {
	var srcArg = ""
	var dstArg = ""
	if len(args) > 0 {
		srcArg = args[0]
	}
	if len(args) > 1 {
		dstArg = args[1]
	}

	c.Engine.Mv(srcArg, dstArg)
}

func (c *MvCommand) Verify(args []string) error {
	if len(args) > 2 {
		return common.NewStringError("to many arguments")
	}
	return nil
}

func (c *MvCommand) CommandString() string {
	return "mv"
}

func (o *MvCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: true}
}
