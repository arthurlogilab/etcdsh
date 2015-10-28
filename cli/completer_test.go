package cli

import "testing"
import (
	"github.com/kamilhark/etcdsh/etcdclient"
	"github.com/kamilhark/etcdsh/commands"
	"github.com/kamilhark/etcdsh/pathresolver"
	"strings"
)

var etcdClient = etcdclient.NewEtcdClientMock()
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

func TestCompleteFirstDirArgumentWhenInRootDir(t *testing.T) {
	rootNode := etcdclient.Node{}
	rootNode.Nodes = []etcdclient.Node{createDirNode("/aa"), createDirNode("/ab"), createDirNode("/bb"), createValueNode("aaa")}

	response := &etcdclient.Response{"", rootNode}
	etcdClient.MockGet(pathResolver.CurrentPath(), response)

	hints := completer("cd a")

	assertLength(t, hints, 2)
	assertContainHint(t, hints, "cd aa", "cd ab")
}

func TestCompleteFirstDirArgumentWhenInChildDir(t *testing.T) {
	pathResolver.Add("child")
	rootNode := etcdclient.Node{}
	rootNode.Nodes = []etcdclient.Node{createDirNode("/child/aa"), createDirNode("/child/ab")}

	response := &etcdclient.Response{"", rootNode}
	etcdClient.MockGet(pathResolver.CurrentPath(), response)

	hints := completer("cd a")

	assertLength(t, hints, 2)
	assertContainHint(t, hints, "cd aa", "cd ab")
}

func assertContainHint(t *testing.T, actualHints []string, expectedHints ...string) {
	for _, hint := range expectedHints {
		found := false
		for _, a := range actualHints {
			if a == hint {
				found = true
			}
		}
		if !found {
			t.Errorf("actual hints [%s] does not contain %s", strings.Join(actualHints, ","), hint)
		}
	}
}

func assertLength(t *testing.T, slice []string, expectedLength int) {
	if len(slice) != expectedLength {
		t.Errorf("expected size %d but was %d", expectedLength, len(slice))
		t.Fail()
	}
}

func createDirNode(key string) (node etcdclient.Node) {
	node.Dir = true
	node.Key = key
	return
}

func createValueNode(key string) (node etcdclient.Node) {
	node.Dir = false
	node.Key = key
	return
}

