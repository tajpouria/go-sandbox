package main

import (
	"io"
	"log"
	"net/http"
)

type flushWriter struct {
	w io.Writer
}

func (fw flushWriter) Write(p []byte) (n int, err error) {
	n, err = fw.w.Write(p)
	if f, ok := fw.w.(http.Flusher); ok {
		f.Flush()
	}
	return
}

func main() {
	srv := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: http.HandlerFunc(echoHandler),
	}

	log.Println("Echo server Listening on 127.0.0.1:8080")
	log.Fatal(srv.ListenAndServeTLS("secret/server.crt", "secret/server.key"))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	io.Copy(flushWriter{w: w}, r.Body)
}
