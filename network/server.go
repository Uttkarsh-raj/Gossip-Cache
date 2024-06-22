package network

import (
	"sync"
	"time"
)

// It acts like a gateway from where the nodes can identify
// each other and then can communicate with each other
// In case of failure of the server the nodes will still be able
// to be in contact with each other and share data, although no
// new nodes will be able to join the network.
type Server struct {
	Nodes map[string]*Node
	Mutex sync.Mutex
}

// New Instance of the server
func NewServer() *Server {
	return &Server{
		Nodes: make(map[string]*Node),
	}
}

// Add a new Node/User to the network and inform about other nodes
func (s *Server) AddAndStartGossip(node *Node) {
	s.AddNode(node)
	// Gossip starts...
	// get all the cache data and then share it with other peers
	go func() {
		for {
			go StartGossip(node)
			time.Sleep(time.Second * 5)
			// println(len(node.Cache.Items))
		}
	}()
}

// Add a new Node/User to the network
func (s *Server) AddNode(node *Node) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	// After a new node is discovered inform all other nodes about it
	_, pres := s.Nodes[node.Addr]
	if !pres {
		s.Nodes[node.Addr] = node
		for _, remoteNode := range s.Nodes {
			if remoteNode.Addr != node.Addr {
				remoteNode.Neighbor = append(remoteNode.Neighbor, node)
				node.Neighbor = append(node.Neighbor, remoteNode)
			}
		}
	}
}
