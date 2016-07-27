package cli

import "flag"
import "fmt"
import "strings"
import "log"
import "github.com/kamilhark/etcdsh/commands"
import "github.com/kamilhark/etcdsh/pathresolver"

import (
	"time"

	"github.com/coreos/etcd/client"
	"github.com/peterh/liner"
)

func Start() {

	var urls = flag.String("urls", "", "etcd urls")
	var url = flag.String("url", "", "etcd url")

	flag.Parse()

	etcdUrls := strings.Split(*urls, ",")
	etcdUrl := *url

	if etcdUrl != "" {
		if *urls != "" {
			log.Fatal("You must enter --url or --urls")
		} else {
			etcdUrls = []string{etcdUrl}
		}
	}

	if len(etcdUrls) == 1 && etcdUrls[0] == "" {
		log.Fatal("You must enter at least one URL. Use --url or --urls")
	}

	pathResolver := new(pathresolver.PathResolver)
	cfg := client.Config{
		Endpoints: etcdUrls,
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	api := client.NewKeysAPI(c)

	fmt.Println("connected to etcd")

	console := liner.NewLiner()
	console.SetTabCompletionStyle(liner.TabCircular)
	commandsArray := []commands.Command{
		&commands.ExitCommand{State: console},
		&commands.CdCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.LsCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.DumpCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.GetCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.SetCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.RmCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.RmDirCommand{PathResolver: pathResolver, KeysApi: api},
		&commands.MkDirCommand{PathResolver: pathResolver, KeysApi: api},
	}

	defer console.Close()
	console.SetCtrlCAborts(true)
	completer := (&Completer{api, commandsArray, pathResolver}).Get
	console.SetCompleter(completer)

	for {
		line, err := console.Prompt(pathResolver.CurrentPath() + "> ")

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

func printPrompt(pathResolver *pathresolver.PathResolver) {
	fmt.Print()
}
