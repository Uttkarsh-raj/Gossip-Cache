package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	testServer := NewServer()
	node1 := CreateNode("192.167.23.32")
	node2 := CreateNode("192.162.43.35")
	node3 := CreateNode("132.147.25.22")

	testServer.Nodes[node1.Addr] = node1
	testServer.Nodes[node2.Addr] = node2
	testServer.Nodes[node3.Addr] = node3

	testNode := CreateNode("123.23.34.54")
	testServer.AddNode(testNode)

	expectedNeighbors := []*Node{node1, node2, node3}

	// Check if testNode has the correct neighbors
	assert.ElementsMatch(t, expectedNeighbors, testNode.Neighbor, "Neighbors do not match")

	// Check if the other nodes have testNode as their neighbor
	for _, node := range expectedNeighbors {
		for i := 0; i < len(node.Neighbor); i++ {
			assert.Contains(t, node.Neighbor[i].Addr, testNode.Addr, "testNode not found in neighbors")
		}
	}
}
