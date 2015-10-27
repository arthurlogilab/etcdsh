package etcdclient

type EtcdClientMock struct {
	mockedGetMethodInvocations map[string]*Response
}

func NewEtcdClientMock() *EtcdClientMock {
	return &EtcdClientMock{mockedGetMethodInvocations:make(map[string]*Response)}
}

func (c *EtcdClientMock) Version() (string, error) {
	return "mock", nil
}

func (c *EtcdClientMock) Get(key string) (*Response, error) {
	return c.mockedGetMethodInvocations[key], nil
}

func (c *EtcdClientMock) Set(key, value string) error {
	return nil
}

func (c *EtcdClientMock) Delete(key string) error {
	return nil
}

func (c * EtcdClientMock) MockGet(key string, response *Response) {
	c.mockedGetMethodInvocations[key] = response
}
