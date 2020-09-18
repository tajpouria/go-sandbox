package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
)

const url = "https://127.0.0.1:8080"

func main() {
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

	client.Transport = &http2.Transport{
		TLSClientConfig: tlsConfig,
	}

	pr, pw := io.Pipe()

	req, err := http.NewRequest(http.MethodPut, url, ioutil.NopCloser(pr))
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Got Response %d\n", res.StatusCode)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Fprintf(pw, "It is not %v\n", time.Now())
		}
	}()

	_, err = io.Copy(os.Stdout, res.Body)
	log.Fatal(err)
}
