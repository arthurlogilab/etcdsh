package etcdclient

import "net/http"
import "bytes"
import "net/url"
import "io/ioutil"
import "encoding/json"
import "github.com/kamilhark/etcdsh/common"

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

func (c *EtcdClient) Get(key string) (*Response, error) {
	resp, err := c.httpClient.Get(c.url + "/v2/keys/" + key)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, common.NewStringError("key not found")
	}
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	etcdResponse := new(Response)
	err = json.Unmarshal(jsonData, etcdResponse)
	return etcdResponse, err
}

func (c *EtcdClient) Set(key, value string) error {
	values := url.Values{}
	values.Set("value", value)
	url := c.url + "/v2/keys" + key
	data := bytes.NewBufferString(values.Encode())
	request, err := http.NewRequest("PUT", url, data)
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return common.NewStringError("key not found")
	}
	return nil
}

func (c *EtcdClient) Delete(key string) error {
	url := c.url + "/v2/keys" + key + "?recursive=true"
	request, err := http.NewRequest("DELETE", url, nil)
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return common.NewStringError("key not found")
	}
	return nil
}
