package cli

import "bufio"
import "encoding/json"
import "flag"
import "fmt"
import "io/ioutil"
import "net/http"
import "os"
import "strings"
import "github.com/kamilhark/etcd-console/commands"
import "github.com/kamilhark/etcd-console/path"

func Start() {
	etcdUrl := getEtcdUrl()
	httpClient := &http.Client{}
	etcdPath := new(path.EtcdPath)

	fetchAndPrintVersion(httpClient, etcdUrl)
	printPrompt(etcdPath)

	commandsArray := [...]commands.Command{
		commands.NewExitCommand(),
		commands.NewCdCommand(etcdPath),
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
				commandHandler.Handle(args)
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

func fetchAndPrintVersion(httpClient *http.Client, etcdUrl *string) {
	resp, err := httpClient.Get(*etcdUrl + "/version")

	if err != nil {
		fmt.Println(err)
		return
	}

	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	var version interface{}

	err = json.Unmarshal(jsonDataFromHttp, &version)
	m := version.(map[string]interface{})

	fmt.Print("Connected to " + (*etcdUrl) + ", version ")
	fmt.Println(m["releaseVersion"])
}
