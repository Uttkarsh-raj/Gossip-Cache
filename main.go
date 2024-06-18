package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Uttkarsh-raj/Dist-Cache/network"
)

func main() {
	server := network.NewServer()

	go simulateAdditionOfNodes(server)

	// Register the node to the server to know other peers and acces the distributed cache
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		node, present := server.Nodes[ip]
		if !present {
			node = network.CreateNode(ip)
		}
		server.AddNode(node)
	})

	// Retrieve the cache from the in-memory cache as the data is stored locally after knowing the peers
	// If not known then will eventually as the nodes are queried randomly
	// The key needs to be provided in the params i.e. /get/key
	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		node, present := server.Nodes[ip]
		if !present {
			http.Error(w, "New nodes need to be registered. Please try to connect using the gateway using the '/' route", http.StatusBadRequest)
			return
		}

		key := strings.TrimPrefix(r.URL.Path, "/get/")
		if key == "" {
			http.Error(w, "Key is required i.e. /get/key", http.StatusBadRequest)
			return
		}

		cacheItem, exists := node.Cache.Items[key]
		if !exists {
			http.Error(w, "Key not found in cache", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(cacheItem)

	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func simulateAdditionOfNodes(server *network.Server) {
	// Simulating addition of new nodes in the network
	testNode := network.CreateNode("123.23.45.33")
	testNode2 := network.CreateNode("232.122.22.10")
	testNode.Cache.Items["key"] = *network.NewCacheItem("key", "hello-World", 100)
	testNode2.Cache.Items["Key"] = *network.NewCacheItem("Key", "new-Key", 100)
	go func() {
		server.Nodes["123.23.45.33"] = testNode
		time.Sleep(time.Duration(time.Second * 3))
	}()
	go func() {
		server.Nodes["232.122.22.10"] = testNode2
		time.Sleep(time.Duration(time.Second * 7))
	}()
}
