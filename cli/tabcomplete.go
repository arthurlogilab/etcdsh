package cli
import (
	"strings"
	"github.com/kamilhark/etcdsh/etcdclient"
	"github.com/kamilhark/etcdsh/commands"
	"github.com/kamilhark/etcdsh/pathresolver"
)

func GetCompleter(etcdClient *etcdclient.EtcdClient, commandsArray []commands.Command, pathResolver *pathresolver.PathResolver) func (line string) (c []string) {

	return func (line string) (c []string) {
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
	}


}
