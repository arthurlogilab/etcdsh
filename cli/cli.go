package cli

import "flag"
import "fmt"
import "strings"
import "log"
import "github.com/kamilhark/etcdsh/etcdclient"
import "github.com/kamilhark/etcdsh/commands"
import "github.com/kamilhark/etcdsh/pathresolver"
import "github.com/peterh/liner"


func Start() {
	etcdUrl := getEtcdUrl()
	pathResolver := new(pathresolver.PathResolver)
	etcdClient := etcdclient.NewEtcdClient(*etcdUrl)

	version, err := etcdClient.Version()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("connected to etcd " + version)

	console := liner.NewLiner()
	console.SetTabCompletionStyle(liner.TabCircular)
	commandsArray := []commands.Command{
		commands.NewExitCommand(console),
		commands.NewCdCommand(pathResolver, etcdClient),
		commands.NewLsCommand(pathResolver, etcdClient),
		commands.NewGetCommand(pathResolver, etcdClient),
		commands.NewSetCommand(pathResolver, etcdClient),
		commands.NewRmCommand(pathResolver, etcdClient),
	}

	defer console.Close()
	console.SetCtrlCAborts(true)
	//todo extract to seperate module and unit test it!
	console.SetCompleter(func(line string) (c []string) {

		tokens := strings.Split(line, " ")

		if len(tokens) == 1 { //user entered only a command (or part of a command) name without arguments
			for _, commandHandler := range commandsArray {
				if strings.HasPrefix(commandHandler.CommandString(), line) {
					c = append(c, commandHandler.CommandString())
				}
			}

		} else if len(tokens) == 2 { //user entered full command name and part of argument


			res, _ := etcdClient.Get(pathResolver.CurrentPath())
			nodes := res.Node.Nodes

			for _, commandHandler := range commandsArray {

				if strings.Trim(line, " ") == commandHandler.CommandString() || strings.HasPrefix(line, commandHandler.CommandString()) {

					for _, node := range nodes {

						if !node.Dir {
							continue
						}

						key := strings.TrimPrefix(strings.TrimPrefix(node.Key, pathResolver.CurrentPath()), "/")

						if strings.HasPrefix(key, tokens[1]) {
							c = append(c, commandHandler.CommandString() + " " + key)
						}
					}

				}
			}

		}

		return
	});

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

func getEtcdUrl() *string {
	var url = flag.String("url", "http://localhost:4001", "etcd url")
	flag.Parse()
	return url
}

func printPrompt(pathResolver *pathresolver.PathResolver) {
	fmt.Print()
}
