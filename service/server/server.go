package server

import (
	"crypto/tls"
	"github.com/maxlambrecht/svid-exercise/service/validator"
	"net/http"
	"log"
	"fmt"
)

type AuthServer struct {
	Addr          string
	CertFile      string
	KeyFile       string
	SpiffeID      string
	CertValidator validator.Validator
}

// Channels to enable sending the server a signal to shutdown
var shutdown = make(chan int)
var done = make(chan int)

func (s *AuthServer) Start() {
	cfg := &tls.Config{
		ClientAuth:         tls.RequireAnyClientCert,
		InsecureSkipVerify: true,
	}
	server := &http.Server{
		Addr:      s.Addr,
		TLSConfig: cfg,
	}

	http.HandleFunc("/auth", s.authenticateHandler)


	// Define a way to send a signal to the server to shutdown
	// Used in integration test
	go func() {
		<-shutdown
		server.Close()
		close(done)
	}()


	fmt.Printf("Server listening on address %s\n", server.Addr)
	if err := server.ListenAndServeTLS(s.CertFile, s.KeyFile); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", s.Addr, err)
	}

	<-done
}

func (s *AuthServer) authenticateHandler(response http.ResponseWriter, request *http.Request) {

	err := s.CertValidator.ValidateID(s.SpiffeID, request.TLS.PeerCertificates[0])
	if err != nil {
		response.WriteHeader(401)
		return
	}
	response.WriteHeader(200)
}
