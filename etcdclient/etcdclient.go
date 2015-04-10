package etcdclient

import "net/http"
import "io/ioutil"
import "encoding/json"

func NewEtcdClient(etcdUrl string) *EtcdClient {
	etcdClient := new(EtcdClient)
	etcdClient.httpClient = http.Client{}
	etcdClient.url = etcdUrl
	return etcdClient
}

type EtcdClient struct {
	url        string
	httpClient http.Client
}

func (c *EtcdClient) Version() (error, string) {
	resp, err := c.httpClient.Get(c.url + "/version")

	if err != nil {
		return err, ""
	}

	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err, ""
	}

	version := Version{}

	err = json.Unmarshal(jsonDataFromHttp, &version)

	return nil, version.String()
}

//func (c *EtcdClient) Get(string key) string {

//}
