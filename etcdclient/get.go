package etcdclient

type Node struct {
	Key           string
	Dir           bool
	Value         string
	Nodes         []Node
	ModifiedIndex int
	CreatedIndex  int
}

type Get struct {
	Action string
	Node   Node
}

func (g *Get) String() string {
	return g.Node.Value
}
