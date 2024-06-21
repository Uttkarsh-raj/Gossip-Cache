package test

import (
	"testing"
	"time"

	"github.com/Uttkarsh-raj/Gossip-Cache/network"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	testServer := network.NewServer()
	node1 := network.CreateNode("192.167.23.32")
	node2 := network.CreateNode("192.162.43.35")
	node3 := network.CreateNode("132.147.25.22")

	testServer.Nodes[node1.Addr] = node1
	testServer.Nodes[node2.Addr] = node2
	testServer.Nodes[node3.Addr] = node3

	testNode := network.CreateNode("123.23.34.54")
	testServer.AddNode(testNode)

	expectedNeighbors := []*network.Node{node1, node2, node3}

	// Check if testNode has the correct neighbors
	assert.ElementsMatch(t, expectedNeighbors, testNode.Neighbor, "Neighbors do not match")

	// Check if the other nodes have testNode as their neighbor
	for _, node := range expectedNeighbors {
		for i := 0; i < len(node.Neighbor); i++ {
			assert.Contains(t, node.Neighbor[i].Addr, testNode.Addr, "testNode not found in neighbors")
		}
	}
}

func TestSimulateAdditionOfNodes(t *testing.T) {
	testServer := network.NewServer()

	simulateAdditionOfNodes(testServer)

	// Verify that nodes have been added to the server
	_, exists1 := testServer.Nodes["123.23.45.33"]
	_, exists2 := testServer.Nodes["232.122.22.10"]

	assert.True(t, exists1, "Node 123.23.45.33 should exist in the server")
	assert.True(t, exists2, "Node 232.122.22.10 should exist in the server")

	// Verify the cache items
	node1 := testServer.Nodes["123.23.45.33"]
	node2 := testServer.Nodes["232.122.22.10"]

	item1, exists1, _ := node1.Cache.Get("key")
	item2, exists2, _ := node2.Cache.Get("Key")

	assert.True(t, exists1, "Key 'key' should exist in node1's cache")
	assert.Equal(t, "hello-World", item1.Value, "Value for 'key' in node1's cache should be 'hello-World'")

	assert.True(t, exists2, "Key 'Key' should exist in node2's cache")
	assert.Equal(t, "new-Key", item2.Value, "Value for 'Key' in node2's cache should be 'new-Key'")
}

func simulateAdditionOfNodes(server *network.Server) {
	// Simulating addition of new nodes in the network
	testNode := network.CreateNode("123.23.45.33")
	testNode2 := network.CreateNode("232.122.22.10")
	testNode.Cache.Add("key", "hello-World", 5000)
	testNode2.Cache.Add("Key", "new-Key", 5000)

	server.AddNode(testNode)
	time.Sleep(time.Second * 3)
	server.AddNode(testNode2)
	time.Sleep(time.Second * 7)
}
