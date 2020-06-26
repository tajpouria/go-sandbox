package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const (
		usingLocalServers = false
	)

	var links = []string{}

	switch usingLocalServers {
	case true:
		links = []string{
			"http://localhost:4000/book",
			"http://localhost:8000/book",
		}
	default:
		links = []string{
			"http://github.com",
			"http://google.com",
			"http://instagram.com",
			"http://golang.org",
			"http://stackoverflow.com",
		}
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(l string, c chan string) {
	time.Sleep(3 * time.Second)
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, "might be down!")
		c <- l
		return
	}

	fmt.Println(l, "is Ok.")
	c <- l
}
