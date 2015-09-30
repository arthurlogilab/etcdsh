package cli

import "flag"
import "fmt"
import "strings"
import "log"
import "github.com/kamilhark/etcdsh/etcdclient"
import "github.com/kamilhark/etcdsh/commands"
import "github.com/kamilhark/etcdsh/pathresolver"
import "github.com/peterh/liner"
import "github.com/aybabtme/rgbterm"


var r, g, b uint8 = 17, 47, 193

func Start() {
	etcdUrl := getEtcdUrl()
	pathResolver := new(pathresolver.PathResolver)
	etcdClient := etcdclient.NewEtcdClient(*etcdUrl)

	version, err := etcdClient.Version()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(rgbterm.FgString("connected to etcd " + version, r, g, b))

	console := liner.NewLiner()
	console.SetTabCompletionStyle(liner.TabPrints)
	commandsArray := [...]commands.Command{
		commands.NewExitCommand(console),
		commands.NewCdCommand(pathResolver, etcdClient),
		commands.NewLsCommand(pathResolver, etcdClient),
		commands.NewGetCommand(pathResolver, etcdClient),
		commands.NewSetCommand(pathResolver, etcdClient),
		commands.NewRmCommand(pathResolver, etcdClient),
	}

	defer console.Close()
	console.SetCtrlCAborts(true)
	console.SetCompleter(func(line string) (c []string) {
		for _, commandHandler := range commandsArray {
			if (strings.HasPrefix(commandHandler.CommandString(), line)) {
				c = append(c, commandHandler.CommandString() + " ")
			}
		}

		return
	});

	for ;; {
		line, err := console.Prompt(rgbterm.FgString("etcdsh:" + pathResolver.CurrentPath() + ">", r, g, b))

		if err != nil && err == liner.ErrPromptAborted {
			return
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

func getEtcdUrl() *string {
	var url = flag.String("url", "http://localhost:4001", "etcd url")
	flag.Parse()
	return url
}

func printPrompt(pathResolver *pathresolver.PathResolver) {
	fmt.Print()
}
