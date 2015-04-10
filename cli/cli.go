package cli

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Start() {
	var url = flag.String("url", "http://localhost:4001", "etcd url")
	flag.Parse()

	httpClient := &http.Client{}

	resp, err := httpClient.Get(*url + "/version")

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

	fmt.Print("Connected to " + (*url) + ", version ")
	fmt.Println(m["releaseVersion"])

	printPrompt()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		textCommand := scanner.Text()
		if strings.HasPrefix(textCommand, "cd") {
			args := strings.Split(textCommand, " ")[:1]
			if len(args) == 0 {

			} else {

			}

		} else if strings.EqualFold(textCommand, "exit") {
			os.Exit(0)
		}
		printPrompt()
	}
}

func printPrompt() {
	fmt.Print("/>")
}
