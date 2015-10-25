package cli

import "testing"
import (
	"github.com/kamilhark/etcdsh/etcdclient"
	"github.com/kamilhark/etcdsh/commands"
	"github.com/kamilhark/etcdsh/pathresolver"
)

var etcdClient = &etcdclient.EtcdClientMock{}
var pathResolver = &pathresolver.PathResolver{}
var commandsArray = []commands.Command{
	&commands.CdCommand{pathResolver, etcdClient},
	&commands.LsCommand{pathResolver, etcdClient},
	&commands.GetCommand{pathResolver, etcdClient},
	&commands.SetCommand{pathResolver, etcdClient},
	&commands.RmCommand{pathResolver, etcdClient},
}
var completer = (&Completer{etcdClient, commandsArray, pathResolver}).Get

func TestCompleteCommandsNames(t *testing.T) {
	assertContainHint(t, completer("c"), "cd")
	assertContainHint(t, completer("s"), "set")
	assertContainHint(t, completer("r"), "rm")
	assertContainHint(t, completer("l"), "ls")

	//when there is no input, all commands should be given
	hints := completer("")
	if len(hints) != len(commandsArray) {
		t.Fail()
	}
}

func assertContainHint(t *testing.T, hints []string, expectedHint string) {
	for _, a := range hints {
		if a == expectedHint {
			return
		}
	}
	t.Fail()
}
