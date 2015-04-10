package path

import "strings"

type EtcdPath struct {
	path []string
}

func (p *EtcdPath) Add(subPath string) {
	p.path = append(p.path, subPath)
}

func (p *EtcdPath) String() string {
	return strings.Join(p.path, "/")
}

func (p *EtcdPath) Clear() {
	p.path = []string{}
}
