package server

import (
	"crypto/tls"
	"github.com/maxlambrecht/svid-exercise/service/validator"
	"log"
	"net/http"
)

// AuthServer defines the configuration options
// for the Authentication Server and is the type
// over which is define the Start method to create
// a HTTPS server
type AuthServer struct {
	Addr     string
	CertFile string
	KeyFile  string
	// SVID SpiffeID that is trusted by the server and will be used
	// to validate the SVID certificate provided by the client
	SpiffeID string
	// Used to validate the SpiffeID in the client SVID certificate
	CertValidator validator.Validator
}

// Channels to enable sending the server a signal to shutdown
// They are not exported since they are only used in tests
var shutdown = make(chan int)
var done = make(chan int)

// Configure an instance of an http.Server and run it
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

	// Defines a way to send a signal to the server to shutdown
	// Used as a workaround to handle the shutdown in integration tests
	// It could be used as a way to gracefully shutdown the server when
	// a 'interrupt' signal is received, but it needs to be improved
	// for handling pending connections before shutting down. That
	// would no be useful in this exercise, thus it's not developed further
	go func() {
		// Waiting for 'shutdown' signal
		<-shutdown
		server.Close()
		// Send 'done' signal
		done <- 1
	}()

	log.Printf("Server listening on address %s\n", server.Addr)
	if err := server.ListenAndServeTLS(s.CertFile, s.KeyFile); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", s.Addr, err)
	}

	//Wait for 'done' signal
	<-done
}

// authenticateHandler handle the authentication requests, validates that the certificated provided by the client
// contains a Subject Alternative Name that matches the SpiffeID the server has been configured with
func (s *AuthServer) authenticateHandler(response http.ResponseWriter, request *http.Request) {

	err := s.CertValidator.ValidateID(s.SpiffeID, request.TLS.PeerCertificates[0])
	if err != nil {
		response.WriteHeader(401)
		return
	}
	response.WriteHeader(200)
}
