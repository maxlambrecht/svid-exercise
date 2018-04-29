package main

import (
	"github.com/maxlambrecht/svid-exercise/service/server"
	"flag"
	"fmt"
)


var addr string
var spiffeID string
var certPath string
var keyPath string

func init() {
	flag.StringVar(&addr, "addr", ":3000", "TCP address to listen on")
	flag.StringVar(&spiffeID, "spiffeid", "", "Spiffe ID")
	flag.StringVar(&certPath,"cert", "", "Client Certificate File")
	flag.StringVar(&keyPath, "key", "", "Client Certificate Key File")
	flag.Parse()
}

func main() {

	if spiffeID == "" {
		fmt.Println("Error: Missing SpiffeID")
		return
	}

	if certPath == "" {
		fmt.Println("Error: Missing service certificate")
		return
	}

	if keyPath == "" {
		fmt.Println("Error: Missing service certificate key ")
		return
	}

	authServer := server.AuthServer{
		Addr:     addr,
		SpiffeID: spiffeID,
		CertFile: certPath,
		KeyFile:  keyPath,
		Validator: server.SvidValidator{},
	}

	authServer.Start()
}
