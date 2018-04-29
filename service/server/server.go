package server

import (
	"crypto/tls"
	"github.com/maxlambrecht/svid-exercise/service/validator"
	"log"
	"net/http"
	"fmt"
)

type AuthServer struct {
	Addr          string
	CertFile      string
	KeyFile       string
	SpiffeID      string
	CertValidator validator.Validator
}

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

	fmt.Printf("Server listening on address %s ", server.Addr)
	log.Fatal(server.ListenAndServeTLS(s.CertFile, s.KeyFile))
}

func (s *AuthServer) authenticateHandler(response http.ResponseWriter, request *http.Request) {

	err := s.CertValidator.ValidateID(s.SpiffeID, request.TLS.PeerCertificates[0])
	if err != nil {
		response.WriteHeader(401)
		return
	}
	response.WriteHeader(200)
}
