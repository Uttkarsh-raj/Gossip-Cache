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
			currCacheItem, present := node.Cache.Items[key]
			if !present {
				node.Cache.Items[key] = cacheItem
			}

			// Check the latest change to the cache
			if currCacheItem.TTL > cacheItem.TTL {
				node.Cache.Items[key] = currCacheItem
				selectedNode.Cache.Items[key] = currCacheItem // Update cache to the remote node sharing cache data with
			} else {
				node.Cache.Items[key] = cacheItem
			}
		}
	}
}
