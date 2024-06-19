package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Uttkarsh-raj/Dist-Cache/network"
)

type HandleError struct {
	Success bool
	Message string
	Data    any
}

func NewErrorHandler(success bool, message string, data any) string {
	a := HandleError{
		Success: success,
		Message: message,
		Data:    data,
	}
	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	return string(out)
}

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
			http.Error(w, NewErrorHandler(false, "New nodes need to be registered. Please try to connect using the gateway using the '/' route", nil), http.StatusBadRequest)
			return
		}

		key := strings.TrimPrefix(r.URL.Path, "/get/")
		if key == "" {
			http.Error(w, NewErrorHandler(false, "Key is required i.e. /get/key", nil), http.StatusBadRequest)
			return
		}

		cacheItem, exists, err := node.Cache.Get(key)
		if !exists {
			http.Error(w, NewErrorHandler(false, err.Error(), nil), http.StatusNotFound)
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
			http.Error(w, NewErrorHandler(false, "New nodes need to be registered. Please try to connect using the gateway using the '/' route", nil), http.StatusBadRequest)
			return
		}

		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, NewErrorHandler(false, "Failed to read request body"+err.Error(), nil), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var item network.CacheItem
		err = json.Unmarshal(body, &item)
		if err != nil {
			http.Error(w, NewErrorHandler(false, "Invalid request payload: "+err.Error(), nil), http.StatusBadRequest)
			return
		}

		message, item := node.Cache.Add(item.Key, item.Value, item.TTL)
		response := make(map[string]interface{})
		response["success"] = true
		response["message"] = message
		response["data"] = item
		json.NewEncoder(w).Encode(response)
	})

}

// Requests:
// Connection Req - curl http://localhost:3000/
// Get Req - curl http://localhost:3000/get/:key
// Post Req - curl.exe -X POST http://localhost:3000/set -d '{\"key\":\"test\",\"value\":\"data\",\"ttl\":2000}' -H "Content-Type: application/json"
