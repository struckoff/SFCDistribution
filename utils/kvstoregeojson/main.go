package main

import (
	"flag"
)

const maxColor = 0xffffff

func main() {
	srv := flag.String("address", "http://localhost:9190", "Address of kvstore node")
	node := flag.String("node", "", "(optional) specify node to extract")
	mode := flag.String("mode", "nodes", "nodes/cid")
	flag.Parse()

	if *mode == "nodes" {
		nodesmode(srv, node)
		return
	}

	if *mode == "cid" {
		cidmode(srv)
		return
	}
}
