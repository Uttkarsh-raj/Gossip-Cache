package p2p

// Representation of the remote node trying to connect
type Peer interface{}

// Manages the communication between the different nodes in the network.
// Uses the TCP Protocol
type Transport interface {
	ListenAndAcceptNodes() error
}
