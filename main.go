package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/Uttkarsh-raj/Dist-Cache/network"
)

func main() {
	server := network.NewServer()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		node := network.CreateNode(ip)
		server.AddNode(node)
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
