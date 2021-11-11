package monitoring

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"strings"
)

type ElasticSearchHealth struct {
    Cluster_name string
    Status string
    Timed_out bool
    Number_of_nodes int
    Number_of_data_nodes int
    Active_primaryShards int
    Active_shards int 
    Relocating_shards int 
    Initializing_shards int
    Unassigned_shards int 
    Delayed_unassigned_shards int 
    Number_of_pending_tasks int 
    Number_of_in_flight_fetch int 
    Task_max_waiting_in_queue_millis int 
    Active_shards_percent_as_number float32
}

func GetStatus(url string) (string, error) {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("monitoring-")

	cleanUrl := strings.TrimPrefix(strings.TrimRight(strings.TrimSpace(url), "/"), "http://")

	resp, err := http.Get("http://" + cleanUrl + "/_cluster/health")
	if err != nil {
		log.Printf("[ERROR] %+v \n", err)
		return "", err
	}
    defer resp.Body.Close()
	
	// We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] %+v \n", err)
		return "", err
	}
	log.Println(string(body))

	// Parse body
    var health ElasticSearchHealth
	perseError := json.Unmarshal(body, &health)
	if perseError != nil {
		log.Printf("[ERROR] %+v \n", perseError)
		return "", perseError
	}

	// Return status
	return health.Status, nil
}
