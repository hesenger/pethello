package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	log.Printf("Starting server on port %d", *port)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "OK 3")
	})
	http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
}
