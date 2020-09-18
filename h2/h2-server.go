package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: "127.0.0.1:8080", Handler: http.HandlerFunc(echoHandler)}

	log.Printf("Serving on https://127.0.0.1:8080")
	log.Fatal(srv.ListenAndServeTLS("secret/server.crt", "secret/server.key"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/2nd" {
		fmt.Println("2nd path called")
		w.Write([]byte("This is 2nd path!"))
		return
	}

	if pusher, ok := w.(http.Pusher); ok {
		if err := pusher.Push("/2nd", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}

	w.Write([]byte("This is root path!"))
}
