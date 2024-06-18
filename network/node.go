package network

// It is a representation of the peers/users on the network
type Node struct {
	Addr     string
	Neighbor []*Node
	Cache    Cache
}

func CreateNode(addr string) *Node {
	return &Node{
		Addr:     addr,
		Neighbor: []*Node{},
		Cache: Cache{
			Items: make(map[string]CacheItem),
		},
	}
}
