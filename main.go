package main

import (
	"fmt"
	"log"
	"net/http"

	// "time"

	"github.com/Uttkarsh-raj/Dist-Cache/network"
	"github.com/Uttkarsh-raj/Dist-Cache/routes"
)

func main() {
	fmt.Println("Server Running....")
	server := network.NewServer()

	// To simulate addition of nodes in the network locally
	// go simulateAdditionOfNodesLocally(server)

	routes.AddRoutes(server)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

// func simulateAdditionOfNodesLocally(server *network.Server) {
// 	// Simulating addition of new nodes in the network
// 	testNode := network.CreateNode("123.23.45.33")
// 	testNode2 := network.CreateNode("232.122.22.10")
// 	go func() {
// 		testNode.Cache.Add("key", "hello-World", 5000)
// 		server.AddNode(testNode)
// 		time.Sleep(time.Duration(time.Second * 3))
// 	}()
// 	go func() {
// 		testNode2.Cache.Add("Key", "new-Key", 5000)
// 		server.AddNode(testNode2)
// 		time.Sleep(time.Duration(time.Second * 7))
// 	}()
// }
