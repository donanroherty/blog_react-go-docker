package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type timeJSON struct {
	ServerTime string `json:"server_time"`
}

const (
	port       = 80
	staticPath = "./static"
)

func main() {
	http.HandleFunc("/api/time", handleGetTime)

	www := http.FileServer(http.Dir(staticPath))
	http.Handle("/", www)

	server := &http.Server{Addr: fmt.Sprint(":", port), Handler: nil}

	fmt.Println("Server listening on port", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("ListenAnServe error: %v", err)
	}
}

func handleGetTime(w http.ResponseWriter, r *http.Request) {
	out := timeJSON{ServerTime: time.Now().String()}

	err := json.NewEncoder(w).Encode(out)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
