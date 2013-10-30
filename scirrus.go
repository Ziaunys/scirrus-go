package scirrus

import (
	"net/http"
	"rand"
)

type ClientConf struct {
	ClientID	string
	MinRange	int
	MaxRange	int
	WorkerLimit	int
	Clusters	int
	Tracks		int
}

type Collector struct {
	
}

func track_id_range(min int, max int, n int) (ids []int) {
	for i := 0; i < n; i++ {
		rand_id := rand.Intn(max) % max
		if rand_id < min {
			rand_id := rand_id + min
		}
		ids = append(ids, rand_id)
	}
}

func get_tracks() {
id_pool = track_id_range(
}
