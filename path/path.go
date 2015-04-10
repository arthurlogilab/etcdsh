package path

import "strings"
import "fmt"

type EtcdPath struct {
	path []string
}

func (p *EtcdPath) Add(subPath string) {
	p.path = append(p.path, subPath)
	fmt.Println(p.path)
}

func (p *EtcdPath) String() string {
	return strings.Join(p.path, "/")
}
