package commands

import (
	"strings"

	"github.com/kamilhark/etcdsh/common"
	"github.com/kamilhark/etcdsh/engine"
)

type CpCommand struct {
	Engine engine.Engine
}

func (c *CpCommand) Supports(command string) bool {
	return strings.EqualFold(command, "cp")
}

func (c *CpCommand) Handle(args []string) {
	var srcArg = ""
	var dstArg = ""
	if len(args) > 0 {
		srcArg = args[0]
	}
	if len(args) > 1 {
		dstArg = args[1]
	}

	c.Engine.Cp(srcArg, dstArg)
}

func (c *CpCommand) Verify(args []string) error {
	if len(args) > 2 {
		return common.NewStringError("to many arguments")
	}
	return nil
}

func (c *CpCommand) CommandString() string {
	return "cp"
}

func (o *CpCommand) GetAutoCompleteConfig() AutoCompleteConfig {
	return AutoCompleteConfig{Available: true, OnlyDirs: true}
}
