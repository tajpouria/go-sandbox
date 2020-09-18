package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

const url = "https://127.0.0.1:8080"

var httpVersion = flag.Int("version", 2, "HTTP version")

func main() {
	flag.Parse()

	client := &http.Client{}

	caCert, err := ioutil.ReadFile("secret/server.crt")

	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	switch *httpVersion {
	case 1:
		client.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}

	case 2:
		client.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s", err)
	}

	fmt.Printf("Got response %d: %s %s\n",
		resp.StatusCode, resp.Proto, string(body))
}
