package etcdclient

import "net/http"
import "io/ioutil"
import "encoding/json"
import "github.com/kamilhark/etcd-console/common"

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

func (c *EtcdClient) Version() (string, error) {
	resp, err := c.httpClient.Get(c.url + "/version")

	if err != nil {
		return "", nil
	}

	jsonData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	version := Version{}

	err = json.Unmarshal(jsonData, &version)

	return version.String(), nil
}

func (c *EtcdClient) Get(key string) (string, error) {
	resp, err := c.httpClient.Get(c.url + "/v2/keys/" + key)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", common.NewStringError("key not found")
	}

	jsonData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	get := Get{}

	err = json.Unmarshal(jsonData, &get)

	return get.String(), nil

}
