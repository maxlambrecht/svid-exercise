package main

import (
	"log"
	"crypto/tls"
	"net/http"
	"fmt"
)

func main() {
	// Make this script configurable by command line

	// Load client cert
	cert, err:= tls.LoadX509KeyPair("certs/client_cert.pem", "certs/client_key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		InsecureSkipVerify:true,
	}

	// Configure the client
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}


	// Perform the request
	resp, err := client.Get("https://localhost:3000/auth")
	if err != nil {
		fmt.Println(err)
	}

    if resp.StatusCode == 200 {
    	fmt.Println("Authentication succeed")
	}
	fmt.Println(resp.StatusCode)
}
