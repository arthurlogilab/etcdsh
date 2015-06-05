package cli

import "flag"
import "fmt"
import "strings"
import "log"
import "os"
import "bufio"
import "github.com/kamilhark/etcdsh/etcdclient"
import "github.com/kamilhark/etcdsh/commands"
import "github.com/kamilhark/etcdsh/pathresolver"

func Start() {
	etcdUrl := getEtcdUrl()
	pathResolver := new(pathresolver.PathResolver)
	etcdClient := etcdclient.NewEtcdClient(*etcdUrl)

	version, err := etcdClient.Version()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to etcd " + version)

	printPrompt(pathResolver)

	commandsArray := [...]commands.Command{
		commands.NewExitCommand(),
		commands.NewCdCommand(pathResolver, etcdClient),
		commands.NewLsCommand(pathResolver, etcdClient),
		commands.NewGetCommand(pathResolver, etcdClient),
		commands.NewSetCommand(pathResolver, etcdClient),
		commands.NewRmCommand(pathResolver, etcdClient),
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Split(line, " ")
		if len(tokens) == 0 {
			continue
		}

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
	fmt.Print(pathResolver.CurrentPath() + ">")
}
