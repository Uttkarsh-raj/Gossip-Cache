package network

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartGossip(t *testing.T) {
	// Create nodes
	node1 := CreateNode("192.167.23.32")
	node2 := CreateNode("192.162.43.35")
	node3 := CreateNode("132.147.25.22")

	// Set up initial cache data
	node1.Cache.Items["key1"] = CacheItem{Key: "key1", Value: "value1", TTL: time.Now().Add(1 * time.Hour).UnixNano()}
	node2.Cache.Items["key2"] = CacheItem{Key: "key2", Value: "value2", TTL: time.Now().Add(2 * time.Hour).UnixNano()}
	node3.Cache.Items["key3"] = CacheItem{Key: "key3", Value: "value3", TTL: time.Now().Add(3 * time.Hour).UnixNano()}

	// Establish neighbors
	testNode := CreateNode("133.17.233.132")
	testNode.Neighbor = append(testNode.Neighbor, node2)
	testNode.Neighbor = append(testNode.Neighbor, node1)
	testNode.Neighbor = append(testNode.Neighbor, node3)

	// Invoke StartGossip on node1
	StartGossip(testNode)

	// Check if caches are correctly shared
	assert.True(t, CheckCacheValues(testNode, node1) || CheckCacheValues(testNode, node2) || CheckCacheValues(testNode, node3), "Gossip failed to share cache correctly")
}

func CheckCacheValues(testNode, selectedNode *Node) bool {
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
