package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Uttkarsh-raj/Gossip-Cache/network"
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
			dissnode, pres := server.DisconnectedNodes[ip]
			if !pres {
				node = network.CreateNode(ip)
			} else {
				server.ReviveNode(dissnode)
				node = dissnode
			}
		}
		// add nodes and retrieve other nodes
		server.AddAndStartGossip(node)

		response := make(map[string]interface{})
		response["success"] = true
		response["message"] = "Successfully connected to server."
		response["data"] = nil

		json.NewEncoder(w).Encode(response)

	})

	// Retrieve the cache from the in-memory cache as the data is stored locally after knowing the peers
	// If not known then will eventually as the nodes are queried randomly
	// The key needs to be provided in the params i.e. /get/key
	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		server.Mutex.Lock()
		defer server.Mutex.Unlock()
		ip := strings.Split(r.RemoteAddr, ":")[0]
		node, present := server.Nodes[ip]
		if !present {
			dissNode, pres := server.DisconnectedNodes[ip]
			if !pres {
				http.Error(w, NewErrorHandler(false, "New nodes need to be registered. Please try to connect using the gateway using the '/' route", nil), http.StatusBadRequest)
				return
			}
			node = dissNode
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

		cacheItem.TTL = cacheItem.TTL - time.Now().UnixMilli()

		response := make(map[string]interface{})
		response["success"] = true
		response["message"] = "Connected to the server successfully!!"
		response["data"] = cacheItem

		json.NewEncoder(w).Encode(response)

	})

	// Set a new value for the key in the cache
	// If the data is already present it will update that value
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		server.Mutex.Lock()
		defer server.Mutex.Unlock()
		ip := strings.Split(r.RemoteAddr, ":")[0]
		node, present := server.Nodes[ip]
		if !present {
			dissNode, pres := server.DisconnectedNodes[ip]
			if !pres {
				http.Error(w, NewErrorHandler(false, "New nodes need to be registered. Please try to connect using the gateway using the '/' route", nil), http.StatusBadRequest)
				return
			}
			node = dissNode
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

	// // Remove the cache from the in-memory cache as the data is stored locally after knowing the peers
	// // The key needs to be provided in the params i.e. /get/key
	// http.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
	// 	server.Mutex.Lock()
	// 	defer server.Mutex.Unlock()
	// 	ip := strings.Split(r.RemoteAddr, ":")[0]
	// 	node, present := server.Nodes[ip]
	// 	if !present {
	// 		dissNode, pres := server.DisconnectedNodes[ip]
	// 		if !pres {
	// 			http.Error(w, NewErrorHandler(false, "New nodes need to be registered. Please try to connect using the gateway using the '/' route", nil), http.StatusBadRequest)
	// 			return
	// 		}
	// 		node = dissNode
	// 	}

	// 	key := strings.TrimPrefix(r.URL.Path, "/delete/")
	// 	if key == "" {
	// 		http.Error(w, NewErrorHandler(false, "Key is required i.e. /get/key", nil), http.StatusBadRequest)
	// 		return
	// 	}

	// 	item, exists, err := node.Cache.Get(key)
	// 	if !exists {
	// 		http.Error(w, NewErrorHandler(false, err.Error(), nil), http.StatusNotFound)
	// 		return
	// 	}

	// 	item.TTL = 0
	// 	node.Cache.Update(key, item)
	// 	fmt.Println(item)

	// 	response := make(map[string]interface{})
	// 	response["success"] = true
	// 	response["message"] = "Connected to the server successfully!!"
	// 	response["data"] = item

	// 	json.NewEncoder(w).Encode(response)

	// })

	// Disconnects from the network and enables you use it as an in-memory cache
	http.HandleFunc("/disconnect", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]

		node, present := server.Nodes[ip]
		if !present {
			http.Error(w, NewErrorHandler(false, "Node not found", nil), http.StatusNotFound)
			return
		}

		node.CancelFunc()
		// Remove the node from the server
		server.RemoveNode(node)

		// Send response after removal is completed
		// response := make(map[string]interface{})
		// response["success"] = true
		// response["message"] = "Node disconnected successfully."
		// response["data"] = node
		// json.NewEncoder(w).Encode(response)
	})

}

// Requests:
// Connection Req - curl http://localhost:3000/
// Get Req - curl http://localhost:3000/get/:key
// Post Req - curl.exe -X POST http://localhost:3000/set -d '{\"key\":\"test\",\"value\":\"data\",\"ttl\":2000}' -H "Content-Type: application/json"
