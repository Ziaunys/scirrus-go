package main

import (
	"fmt"
	"flag"
    client "github.com/Ziaunys/scirrus-go/client"
    kmeans "github.com/Ziaunys/scirrus-go/kmeans"
)

func main() {
	var ver = "0.0.1"
	c := new(client.Client)
    c.ClientID = "6482e060ad0be4c70ee8cf6df6ff7aeb"

	fmt.Printf("Scirrus v%s\n", ver)

    c.MinRange = 100
    c.MaxRange = 50000
    c.Tracks = *flag.Int("-t", 5, "the number of tracks to analyze")
    c.Clusters = *flag.Int("-k", 3, "the number of desired clusters")
    c.WorkerLimit = *flag.Int("-w", 100, "the number of workers to use")
    c.MinComments = *flag.Int("-n", 10, "min required comments per track")
	flag.Parse()
    timestamps := c.GetCommentsTs()
    for k, v := range timestamps {
        fmt.Printf("\nTrack ID: %v, Comment Timestamps: %v\n", k, v)
    }
}
