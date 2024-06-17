package network

import "sync"

// It acts like a gateway from where the nodes can identify
// each other and then can communicate with each other
// In case of failure of the server the nodes will still be able
// to be in contact with each other and share data, although no
// new nodes will be able to join the network.
type Server struct {
	Nodes map[string]*Node
	Mutex sync.Mutex
}

func NewServer() *Server {
	return &Server{
		Nodes: make(map[string]*Node),
	}
}

func (s *Server) AddNode(node *Node) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, pres := s.Nodes[node.Addr]
	if !pres {
		s.Nodes[node.Addr] = node
		for _, remoteNode := range s.Nodes {
			if remoteNode.Addr != node.Addr {
				remoteNode.Neighbor = append(remoteNode.Neighbor, *node)
				node.Neighbor = append(node.Neighbor, *remoteNode)
			}
		}
		// After a new node inform all other nodes about it too
		// Gossip starts...
		// get all the cache data and then share it with other peers
	}
}
