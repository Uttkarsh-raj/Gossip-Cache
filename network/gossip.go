package network

import (
	"math/rand"
	"time"
)

func StartGossip(node *Node) {
	// select a random neighbor if neighbors are present
	if len(node.Neighbor) > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		nodeIndex := r.Intn(len(node.Neighbor))
		selectedNode := node.Neighbor[nodeIndex]

		// share cache/data with them and add there cache/data to here
		for key, cacheItem := range selectedNode.Cache.Items {
			// if expired no need to do anything
			if cacheItem.TTL < time.Now().UnixMicro() {
				selectedNode.Cache.Delete(key)
				continue
			}
			currCacheItem, present := node.Cache.Items[key]
			if !present {
				node.Cache.Add(key, cacheItem.Value, cacheItem.TTL-time.Now().UnixNano())
			}
			if currCacheItem.TTL > cacheItem.TTL {
				// Check the latest change to the cache
				// Update cache to the remote node sharing cache data with
				selectedNode.Cache.Update(key, currCacheItem)
			} else {
				node.Cache.Update(key, cacheItem)
			}
		}
	}
}
