package main

import (
	"github.com/maxlambrecht/scytale-exercice/service/server"
	"path/filepath"
)

func main() {

	// Make this script configurable by command line

	authServer := server.AuthServer{
		Addr:     ":3000",
		SpiffeID: "spiffe://example.com/service",
		CertFile: filepath.Join("certs", "server_cert.pem"),
		KeyFile:  filepath.Join("certs", "server_key.pem"),
		Validator: server.CertValidator{},
	}

	authServer.Start()
}
