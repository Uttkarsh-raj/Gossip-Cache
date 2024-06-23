package network

import "context"

// It is a representation of the peers/users on the network
type Node struct {
	Addr       string
	Neighbor   []*Node
	Cache      Cache
	Ctx        context.Context
	CancelFunc context.CancelFunc
}

func CreateNode(addr string) *Node {
	context, cancel := context.WithCancel(context.Background())
	return &Node{
		Addr:     addr,
		Neighbor: []*Node{},
		Cache: Cache{
			Items: make(map[string]CacheItem),
		},
		Ctx:        context,
		CancelFunc: cancel,
	}
}
