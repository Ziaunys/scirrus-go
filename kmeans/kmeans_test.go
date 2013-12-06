package kmeans

import (
    "fmt"
    "io/ioutil"
    "math"
    "math/rand"
    "strings"
    "strconv"
    "testing"
)

func RandRange(min, max float64) float64 {
    rand_id := rand.Intn(max) % max
    if rand_id < int(min) {
        rand_id = rand_id + int(min)
        }
    return float64(rand_id)
}

func GenerateClusteredData(num_points, k int, max_val, max_dist float64) (clustered []float64) {
    var (
     cluster_seeds []float64
     next_point float64
     match bool
    )
    cluster_seeds = append(cluster_seeds, float64(rand.Intn(max)))
    for i := 0; i < k-1; i++ {
        for {
            next_point = float64(rand.Intn(max))
            for _, s := range cluster_seeds {
               if math.Abs(s - next_point) > min_dist {
                   match = true
               } else {
                   match = false
                   break
               }

            }
            if match == true {
                cluster_seeds = append(cluster_seeds, next_point)
                break
            }

        }
    }

    for i := 0; i < num_points; i++ {
        center := cluster_seeds[rand.Intn(len(cluster_seeds))]
        dist := float64(rand.Intn(int(max_dist)+1)
        clustered = append(clustered, RandRange((center - dist),(center + dist)))
    }
}

func TestKmeans(t *testing.T) {
    var data []float64
    var p float64
    b, err := ioutil.ReadFile("test_data.csv")
    if err != nil {
        fmt.Println("Could not read file")
    }
    s := string(b[:(len(b) -1)])
    for _, v := range strings.Split(s, ",") {
        p, err = strconv.ParseFloat(v, 64)
        if err != nil {
            fmt.Println("Could not convert string to float.")
        }
        data = append(data, p)
    }
    t.Log("Test data: ", data)
}
