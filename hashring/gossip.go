package hashring

import (
	"sync"

	"gihub.com/Uttkarsh-raj/Dist-Cache/cache"
)

type GossipNode struct {
	Id        string
	Addr      string
	Cache     *cache.Cache
	Neighbors []Node
	Mutex     sync.Mutex
}

func NewGossipNode(id, addr string, cache *cache.Cache) *GossipNode {
	return &GossipNode{
		Id:    id,
		Addr:  addr,
		Cache: cache,
	}
}

// Add Neighbors to the Gossip-Node
func (g *GossipNode) AddNeighbors(node Node) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	g.Neighbors = append(g.Neighbors, node)
}
