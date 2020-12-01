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
	staticPath = "/srv/www/app"
)

func main() {
	www := http.FileServer(http.Dir(staticPath))
	http.Handle("/", www)

	http.HandleFunc("/api/time", handleGetTime)

	server := &http.Server{Addr: fmt.Sprint(":", port), Handler: nil}

	fmt.Println("Server listening on port", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}

func handleGetTime(w http.ResponseWriter, r *http.Request) {
	out := timeJSON{ServerTime: time.Now().Format("15:04:05 Mon Jan 2 2006")}

	err := json.NewEncoder(w).Encode(out)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
