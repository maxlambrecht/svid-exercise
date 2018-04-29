package server

import (
	"net/http"
	"crypto/tls"
	"github.com/spiffe/go-spiffe/spiffe"
	"log"
	"crypto/x509"
)

type AuthServer struct {
	Addr     string
	CertFile string
	KeyFile  string
	SpiffeID string
	Validator Validator
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

	log.Fatal(server.ListenAndServeTLS(s.CertFile, s.KeyFile))
}

func (s *AuthServer) authenticateHandler(response http.ResponseWriter, request *http.Request) {

	err := s.Validator.ValidateID(s.SpiffeID, request.TLS.PeerCertificates[0])
	if err != nil {
		response.WriteHeader(401)
		return
	}
	response.WriteHeader(200)
}

// Move to another file and add Tests
type Validator interface {
	ValidateID(id string, cert *x509.Certificate) error
}

type CertValidator struct {}

func (v CertValidator) ValidateID(id string, cert *x509.Certificate) error {
	return spiffe.MatchID([]string{id}, cert)
}


