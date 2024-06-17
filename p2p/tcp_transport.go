package p2p

import (
	"fmt"
	"net"
	"sync"
)

// Representation of a remote node in a TCP network
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn net.Conn
	// if the peer is dial and retreive the conn -> outbound: true
	// if the peer accept and retreive the conn -> outbound: false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAdd string
	listener  net.Listener

	mutex sync.Mutex
	peers map[string]Peer
}

func NewTCPTransport(listenAdd string) *TCPTransport {
	return &TCPTransport{
		listenAdd: listenAdd,
		peers:     make(map[string]Peer),
	}
}

func (t *TCPTransport) ListenAndAcceptNodes() error {
	listener, err := net.Listen("tcp", t.listenAdd)
	if err != nil {
		return err
	}
	t.listener = listener
	t.startAccepting()
	return nil
}

func (t *TCPTransport) startAccepting() {
	conn, err := t.listener.Accept()
	if err != nil {
		fmt.Printf("TCP error: %s\n", err)
	}
	go t.handleConn(conn)
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("New connection %v\n", peer)
}
