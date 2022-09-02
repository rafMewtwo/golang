package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", helloWorld).Methods("GET")

	fmt.Println("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Header().Set("Content=Type", "application/json")
		json.NewEncoder(w).Encode("Hello World")
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
