package network

// It is a representation of the peers/users on the network
type Node struct {
	Addr     string
	Neighbor []Node
}

func CreateNode(addr string) *Node {
	return &Node{
		Addr: addr,
	}
}
