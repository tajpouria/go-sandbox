package main

import (
	"fmt"
	"net/http"
)

func main() {
	_, err := http.Get("https://0.0.0.0:8080")

	if err != nil {
		fmt.Println(err)
	}
}
