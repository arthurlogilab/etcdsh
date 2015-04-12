package cli

import "flag"
import "fmt"
import "strings"
import "log"
import "os"
import "bufio"
import "github.com/kamilhark/etcd-console/etcdclient"
import "github.com/kamilhark/etcd-console/commands"
import "github.com/kamilhark/etcd-console/path"

func Start() {
	etcdUrl := getEtcdUrl()
	etcdPath := new(path.EtcdPath)
	etcdClient := etcdclient.NewEtcdClient(*etcdUrl)

	version, err := etcdClient.Version()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to etcd " + version)

	printPrompt(etcdPath)

	commandsArray := [...]commands.Command{
		commands.NewExitCommand(),
		commands.NewCdCommand(etcdPath, etcdClient),
		commands.NewLsCommand(etcdPath, etcdClient),
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

		for _, commandHandler := range commandsArray {
			if commandHandler.Supports(command) {
				err := commandHandler.Verify(args)
				if err != nil {
					fmt.Println(err)
				} else {
					commandHandler.Handle(args)
				}
				break
			}
		}
		printPrompt(etcdPath)
	}
}

func getEtcdUrl() *string {
	var url = flag.String("url", "http://localhost:4001", "etcd url")
	flag.Parse()
	return url
}

func printPrompt(etcdPath *path.EtcdPath) {
	fmt.Print(etcdPath.String() + ">")
}
