package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func readAndWrite() {
	resp, err := http.Get("http://example.com/")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	fmt.Println()
}

func (logWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))

	fmt.Printf("Logging %v bytes", len(p))

	return len(p), nil
}
