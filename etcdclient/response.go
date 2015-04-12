package etcdclient

import "fmt"

type Node struct {
	Key           string
	Dir           bool
	Value         string
	Nodes         []Node
	ModifiedIndex int
	CreatedIndex  int
}

type Response struct {
	Action string
	Node   Node
}

func (g *Response) String() string {
	return fmt.Sprintf("%s %s", g.Node.Key, g.Node.Value)
}
