package cli

import (
	"strings"
	"github.com/kamilhark/etcdsh/etcdclient"
	"github.com/kamilhark/etcdsh/commands"
	"github.com/kamilhark/etcdsh/pathresolver"
)

type Completer struct {
	EtcdClient etcdclient.EtcdClient
	CommandsArray []commands.Command
	PathResolver *pathresolver.PathResolver
}

func (c *Completer) Get(line string) []string {

	tokens := strings.Split(line, " ")

	if len(tokens) == 1 { //user entered only a command (or part of a command) name without arguments
		return c.completeCommand(tokens)
	}

	if len(tokens) == 2 { //user entered full command name and part of argument
		return c.completeArgument(tokens)

	}

	return []string{}
}


func (c *Completer) completeCommand(tokens []string) (result []string) {
	for _, commandHandler := range c.CommandsArray {
		if strings.HasPrefix(commandHandler.CommandString(), tokens[0]) {
			result = append(result, commandHandler.CommandString())
		}
	}
	return
}

func (c *Completer) completeArgument(tokens []string) (result []string) {
	return
}