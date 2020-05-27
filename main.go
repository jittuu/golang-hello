package main

import (
	"encoding/json"
	"net/http"
)

type Hello struct {
	Title string `json:"title"`
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	h := &Hello{Title: "Hello"}
	err := json.NewEncoder(w).Encode(h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
