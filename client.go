package main

import (
	"fmt"
	"flag"
	"scirrus"
)

// Stores the default configuration for the client
func main() {
	var ver = "0.0.1"
	var conf.ClientID = "6482e060ad0be4c70ee8cf6df6ff7aeb"

	conf := new(ClientConf)

	fmt.Printf("Scirrus v%s\n", ver)

	conf.MinRange = 100
	conf.MaxRange = 50000
	conf.Clusters = flag.Int("-k", 3, "the number of desired clusters")
	conf.WorkerLimit = flag.Int("-w", 100, "the number of workers to use")
	conf.Tracks = flag.Int("-n", 10, "min required comments per track")
	flag.Parse()
}
