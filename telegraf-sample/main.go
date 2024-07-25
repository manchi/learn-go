package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type SystemStatus uint

const (
	RUNNING SystemStatus = 1
	FAILING SystemStatus = 0
)

// MetricsResponse is the response that will be sent back to Telegraf
type MetricsResponse struct {
	SystemStatus SystemStatus `json:"systemStatus"`
	TaskStatus   SystemStatus `json:"taskStatus"`
	TaskID       int          `json:"taskID"`
	SystemName   string       `json:"systemName"`
}

// MockSystem to query metrics from
type MockSystem struct {
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}
type Task struct {
	ID int
}

var mockSystems []MockSystem

func main() {
	// init will be called first
	fmt.Println("in main() - start")
	http.HandleFunc("/internal/metrics", fetchMetrics)
	log.Fatal("error::", http.ListenAndServe(":8010", nil))
}

// simple way to get different statuses
func getStatus() SystemStatus {
	if time.Now().UnixNano()%2 == 0 {
		return RUNNING
	}
	return FAILING
}

func init() {
	fmt.Println("in init() - start")
	// used to create 3 mock systems, each with two tasks
	for i := 0; i < 3; i++ {
		mockSystems = append(mockSystems, MockSystem{
			Name: fmt.Sprintf("system-%d", i+1),
			Tasks: []Task{
				Task{ID: 0},
				Task{ID: 1},
			},
		})
	}
	fmt.Println("mockSystems", mockSystems)
	fmt.Println("in init() - end")
}

func fetchMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := []MetricsResponse{}
	// goes through each mock system and collects whether it and its tasks are running
	for _, mockSystem := range mockSystems {
		systemStatus := getStatus()
		for _, task := range mockSystem.Tasks {
			taskStatus := getStatus()
			metrics = append(metrics, MetricsResponse{
				SystemName:   mockSystem.Name,
				SystemStatus: systemStatus,
				TaskID:       task.ID,
				TaskStatus:   taskStatus,
			})
		}
	}
	// create and send json response
	resp, err := json.Marshal(metrics)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
	}
	log.Printf("200 GET /internal/metrics")
	w.Write(resp)

}
