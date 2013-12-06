package kmeans

import (
    "math"
    "math/rand"
    "sort"
)

type Cluster struct {
    Centroid float64
    Set     map[float64]bool
}

func Kmeans(k int, data []float64) {
    var clusters []*Cluster
    var changed bool
    sort.Float64s(data)
    min := data[0]
    max := data[len(data)-1]
    p_size := (max - min)/float64(k)
    for i := min + p_size; i < max; i += p_size {
        clusters = append(clusters, &Cluster{Centroid: i})
    }
    //Initialize the data into random clusters
    for _, d := range data {
    clusters[rand.Intn((len(clusters)-1))].Set[d] = true
    }
    for {
        changed = assign_to_cluster(clusters)
        if changed == true {
            update_centroids(clusters)
        } else {
            break
        }
    }
}

func assign_to_cluster(clusters []*Cluster) (changed bool) {
    var min *Cluster
    changed = false
    for _, c := range clusters {
        for k, _ := range c.Set {
            min  = min_distance(k, clusters)
            _, ok := min.Set[k]
            if ok != true {
                changed = true
            }
            delete(c.Set, k)
       }
    }
    return changed
}

func min_distance(v float64, clusters []*Cluster) *Cluster {
    var dist float64
    min := clusters[len(clusters)-1]
    min_dist := math.Abs(min.Centroid - v)
    for _, c := range clusters {
        dist = math.Abs(c.Centroid - v)
        if dist < min_dist {
            min_dist = dist
            min = c
        }
    }
    return min
}

func update_centroids(clusters []*Cluster) {
    var sum float64
    for _, c := range clusters {
        for k, _ := range c.Set {
            sum += k
        }
        c.Centroid = sum/float64(len(c.Set))
    }
}



