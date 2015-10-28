package cli

import "flag"
import "fmt"
import "strings"
import "log"
import "github.com/kamilhark/etcdsh/etcdclient"
import "github.com/kamilhark/etcdsh/commands"
import "github.com/kamilhark/etcdsh/pathresolver"
import (
	"github.com/peterh/liner"
	"net/http"
)


func Start() {
	etcdUrl := getEtcdUrl()
	pathResolver := new(pathresolver.PathResolver)
	etcdClient := &etcdclient.EtcdClientImpl{etcdUrl, http.Client{}}

	version, err := etcdClient.Version()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("connected to etcd " + version)

	console := liner.NewLiner()
	console.SetTabCompletionStyle(liner.TabCircular)
	commandsArray := []commands.Command{
		&commands.ExitCommand{State: console},
		&commands.CdCommand{PathResolver: pathResolver, EtcdClient: etcdClient},
		&commands.LsCommand{PathResolver: pathResolver, EtcdClient: etcdClient},
		&commands.GetCommand{PathResolver: pathResolver, EtcdClient: etcdClient},
		&commands.SetCommand{PathResolver: pathResolver, EtcdClient: etcdClient},
		&commands.RmCommand{PathResolver: pathResolver, EtcdClient: etcdClient},
	}

	defer console.Close()
	console.SetCtrlCAborts(true)
	completer := (&Completer{etcdClient, commandsArray, pathResolver}).Get
	console.SetCompleter(completer)

	for {
		line, err := console.Prompt(pathResolver.CurrentPath() + ">")

		if err != nil && err == liner.ErrPromptAborted {
			return
		}

		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")
		if len(tokens) == 0 {
			continue
		}

		console.AppendHistory(line)

		command := tokens[0]
		args := tokens[1:]
		found := false
		for _, commandHandler := range commandsArray {
			if commandHandler.Supports(command) {
				found = true
				err := commandHandler.Verify(args)
				if err != nil {
					fmt.Println(err)
				} else {
					commandHandler.Handle(args)
				}
				break
			}
		}
		if !found {
			fmt.Println("invalid command")
		}
		printPrompt(pathResolver)
	}
}

func getEtcdUrl() string {
	var url = flag.String("url", "http://localhost:4001", "etcd url")
	flag.Parse()
	return *url
}

func printPrompt(pathResolver *pathresolver.PathResolver) {
	fmt.Print()
}
