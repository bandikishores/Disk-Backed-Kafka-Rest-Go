package dependency

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Response Generic response for APIs
type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    interface{} `json:"message"`
	Error      string      `json:"error"`
}

var (
	succResponse = Response{StatusCode: 1, Message: "Ok", Error: ""}
)

// AddRouter adds route to work with
func AddRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", getHealthAPI).Methods("GET")
	router.HandleFunc("/sleep", sleepAPI).Methods("GET")
	router.HandleFunc("/postKafka", postKafkaAPI).Methods("GET")
	return router
}

// GetHealthAPI Health Check API
func getHealthAPI(w http.ResponseWriter, req *http.Request) {
	log.Println("Health API Hit")
	json.NewEncoder(w).Encode(&succResponse)
}

// sleepAPI Post Event to the server
func sleepAPI(w http.ResponseWriter, req *http.Request) {
	log.Printf("PreSleep : %s", req.Body)
	time.Sleep(3 * time.Second)
	log.Printf("PostSleep : %s", req.Body)
	json.NewEncoder(w).Encode(&Response{StatusCode: 1, Message: "Sleep Successfully"})
}

// sleepAPI Post Event to the server
func postKafkaAPI(w http.ResponseWriter, req *http.Request) {
	sendMessage("1233", fmt.Sprintf("%s", "test-kafka-go"))
	log.Printf("Succesfully posted message to kafka ")
	json.NewEncoder(w).Encode(&Response{StatusCode: 1, Message: "Message Successfully Received"})
}
