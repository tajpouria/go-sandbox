package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(handler)}

	log.Printf("Serving on https://0.0.0.0:8080")
	log.Fatal(srv.ListenAndServeTLS("secret/server.crt", "secret/server.key"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Proto))
}
