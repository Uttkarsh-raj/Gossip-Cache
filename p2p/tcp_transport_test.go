package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAdd := ":4000"
	tcpTransport := NewTCPTransport(listenAdd)
	assert.Equal(t, tcpTransport.listenAdd, listenAdd)
	assert.Equal(t, tcpTransport.ListenAndAcceptNodes(), nil)
}
