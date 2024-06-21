package test

import (
	"testing"

	"github.com/Uttkarsh-raj/Gossip-Cache/network"
	"github.com/stretchr/testify/assert"
)

func TestStartGossip(t *testing.T) {
	// Create nodes
	node1 := network.CreateNode("192.167.23.32")
	node2 := network.CreateNode("192.162.43.35")
	node3 := network.CreateNode("132.147.25.22")

	// Set up initial cache data
	node1.Cache.Add("key1", "value1", 700)
	node2.Cache.Add("key2", "value2", 700)
	node3.Cache.Add("key3", "value3", 700)

	// Establish neighbors
	testNode := network.CreateNode("133.17.233.132")
	testNode.Neighbor = append(testNode.Neighbor, node2)
	testNode.Neighbor = append(testNode.Neighbor, node1)
	testNode.Neighbor = append(testNode.Neighbor, node3)

	// Invoke StartGossip on node1
	network.StartGossip(testNode)

	// Check if caches are correctly shared
	assert.True(t, CheckCacheValues(testNode, node1) || CheckCacheValues(testNode, node2) || CheckCacheValues(testNode, node3), "Gossip failed to share cache correctly")
}

func CheckCacheValues(testNode, selectedNode *network.Node) bool {
	if len(testNode.Cache.Items) == len(selectedNode.Cache.Items) {
		for key, item := range testNode.Cache.Items {
			val, pres := selectedNode.Cache.Items[key]
			if !pres {
				return false
			}
			if val != item {
				return false
			}
		}
		return true
	}
	return false
}
