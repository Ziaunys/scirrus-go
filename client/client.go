package client

import (
    "errors"
    "io/ioutil"
    "log"
	"net/http"
	"math/rand"
    "strconv"
    "encoding/json"
)

type Client struct {
	ClientID	string
	MinRange	int
	MaxRange	int
    MinComments  int
	WorkerLimit	int
	Clusters	int
	Tracks		int
}

const (
    sc_url = "http://api.soundcloud.com"
)

func (c *Client) track_id_range(min int, max int, n int) (ids []int ) {
    i := 0
    var obj map[string]interface{}
    // Rely on uniqueness of map keys for track ids.
    id_map := make(map[int]bool)
	for {
        rand_id := rand.Intn(max) % max
            if rand_id < min {
                rand_id = rand_id + min
            }

        resp, err := http.Get(sc_url + "/tracks/" + strconv.Itoa(rand_id) + ".json" +
        "?client_id=" + c.ClientID)
        if err != nil {
            log.Println("Error %v",err)
            err.Error()
        }
        if resp.StatusCode != 200 {
            continue
        }
        bod, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Println("Error", err)
            err.Error()
        }
        resp.Body.Close()
        json.Unmarshal(bod, &obj)
        if  obj["comment_count"].(float64) <= float64(c.MinComments) {
            continue
        } else {
            id_map[rand_id] = true
            i++
        }
        if i == n {
            break
        }
	}
    for k, _ := range id_map {
        ids = append(ids, k)
    }
    return ids
}

func (c *Client) GetCommentsTs() (map[int][]float64) {
    var obj []map[string]interface{}
    comms := make(map[int][]float64)
    var ts_list []float64
    id_pool := c.track_id_range(c.MinRange, c.MaxRange, c.Tracks)
    for _, v := range id_pool {
        resp, err := http.Get(sc_url + "/tracks/" + strconv.Itoa(v) +
        "/comments.json" + "?client_id=" + c.ClientID)
        if err != nil {
            log.Println("Error %v %v", resp.Request.RequestURI, err)
            err.Error()
        }
        bod, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Println("Error", err)
            err.Error()
        }
        resp.Body.Close()
        json.Unmarshal(bod, &obj)
        for _, t := range obj {
            ts, ok := t["timestamp"]
            if ok == false {
                err := errors.New("lookup failed for timestamp.")
                log.Println("Key error: ", err)
                err.Error()
            } else if ts == nil {
                log.Println("Warning: timestamp for comment is nil.")
            } else {
                ts_list = append(ts_list, ts.(float64))
            }
        }
        log.Printf("%v", ts_list)
        comms[v] = ts_list
        ts_list = nil
    }
    return comms
}


