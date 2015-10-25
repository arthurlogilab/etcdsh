package etcdclient

type EtcdClientMock struct {
}

func (c *EtcdClientMock) Version() (string, error) {
	return "mock", nil
}

func (c *EtcdClientMock) Get(key string) (*Response, error) {
	return nil, nil
}

func (c *EtcdClientMock) Set(key, value string) error {
	return nil
}

func (c *EtcdClientMock) Delete(key string) error {
	return nil
}
