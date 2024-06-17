package main

import (
	"log"

	"gihub.com/Uttkarsh-raj/Dist-Cache/p2p"
)

func main() {

	tcp := p2p.NewTCPTransport(":3000")
	err := tcp.ListenAndAcceptNodes()
	if err != nil {
		log.Fatal(err)
	}
	select {}
	// log.Fatal(http.ListenAndServe(":3000", nil))
}

// Steps:
// 1) start server
// 2) routes to /Join , Set , Get
// 3) Join -> new node, add to the servers as a node
// 4) and get the data using Gossip protocol
// 5) Set -> set the new data and push to other nodes.
// 6) Get -> wiil have a copy for the value and get it from there....
