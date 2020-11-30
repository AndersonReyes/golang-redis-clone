package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	// "strings"
)

type SetRequest struct {
	Key   string
	Value string
}

type GetRequest struct {
	Key string
}

type Database struct {
	mu     sync.Mutex
	Values map[string]string
}

var db Database

func PingPong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG\r\n")
}

func Set(w http.ResponseWriter, r *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var body SetRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Values[body.Key] = body.Value
	fmt.Fprintf(w, "Set: %s\r\n", body.Key)
}

func Get(w http.ResponseWriter, r *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var body GetRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := db.Values[body.Key]
	fmt.Fprintf(w, "Get: %s\r\n", res)
}

func SetUpRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", PingPong)
	mux.HandleFunc("/set", Set)
	mux.HandleFunc("/get", Get)
}

func InitDb() {
	db = Database{Values: make(map[string]string)}
}

func main() {
	InitDb()
	fmt.Println(db.Values)
	mux := http.NewServeMux()
	SetUpRoutes(mux)
	err := http.ListenAndServe(":7777", mux)
	log.Fatal(err)
}
