package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var url string
var certPath string
var keyPath string

func init() {
	flag.StringVar(&url, "url", "https://localhost:3000/auth", "Authentication service url")
	flag.StringVar(&certPath, "cert", "", "Client Certificate File")
	flag.StringVar(&keyPath, "key", "", "Client Certificate Key File")
	flag.Parse()
}

func main() {

	if certPath == "" {
		fmt.Println("Error: Missing client certificate")
		return
	}

	if keyPath == "" {
		fmt.Println("Error: Missing client certificate key ")
		return
	}

	// Load client certificate
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	// Configure the client
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Perform the request
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response code: %d\n", resp.StatusCode)
	switch resp.StatusCode {
	case 200:
		fmt.Println("Authentication Succeed")
	case 401:
		fmt.Println("Authentication Failed. Invalid SpiffeID")
	default:
		fmt.Println("Authentication Failed. Error not identified")

	}
}
