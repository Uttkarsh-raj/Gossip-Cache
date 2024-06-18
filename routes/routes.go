package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Uttkarsh-raj/Dist-Cache/network"
)

func AddRoutes(server *network.Server) {

	// Register the node to the server to know other peers and acces the distributed cache
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		node, present := server.Nodes[ip]
		if !present {
			node = network.CreateNode(ip)
		}
		// add nodes and retrieve other nodes
		server.AddAndStartGossip(node)

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

		cacheItem, exists, err := node.Cache.Get(key)
		if !exists {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(cacheItem)

	})

	// Set a new value for the key in the cache
	// If the data is already present it will update that value
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {

		ip := strings.Split(r.RemoteAddr, ":")[0]
		node, present := server.Nodes[ip]
		if !present {
			http.Error(w, "New nodes need to be registered. Please try to connect using the gateway using the '/' route", http.StatusBadRequest)
			return
		}

		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var item network.CacheItem
		err = json.Unmarshal(body, &item)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		message, item := node.Cache.Add(item.Key, item.Value, item.TTL)
		response := make(map[string]interface{})
		response["message"] = message
		response["data"] = item
		json.NewEncoder(w).Encode(response)
	})

}

// Requests:
// Connection Req - curl http://localhost:3000/
// Get Req - curl http://localhost:3000/get/:key
// Post Req - curl.exe -X POST http://localhost:3000/set -d '{\"key\":\"test\",\"value\":\"data\",\"ttl\":2000}' -H "Content-Type: application/json"
