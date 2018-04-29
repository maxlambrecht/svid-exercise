package server

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"crypto/tls"
	"crypto/x509"
	"path"
)

func TestAuthenticateHandler(t *testing.T) {

	var assetsDir = "../test-assets"

	testCases := []struct {
		name         string
		spiffeID     string
		expectedCode int
	}{
		{"valid ID", "spiffe://example.com/service", 200},
		{"invalid ID", "spiffe://example.com/other-service", 401},
	}

	authServer := AuthServer{
		// Could be mocked
		Validator: CertValidator{},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			// Configure the service with the SPIFFE ID
			authServer.SpiffeID = tc.spiffeID

			// Load client SVID certificate
			cert, _:= tls.LoadX509KeyPair(path.Join(assetsDir, "cert.pem"),
										  path.Join(assetsDir, "key.pem"))
			certificate := []tls.Certificate{cert}
			x509Cert, _ := x509.ParseCertificate(certificate[0].Certificate[0])

			req, _ := http.NewRequest("GET", "/auth", nil)

			// Configure request TLS with the client SVID
			req.TLS = &tls.ConnectionState{}
			req.TLS.PeerCertificates = make([]*x509.Certificate, 1)
			req.TLS.PeerCertificates[0] = x509Cert

			res := httptest.NewRecorder()

			// Call the function to test
			authServer.authenticateHandler(res, req)

			// Check the expectation
			if res.Code != tc.expectedCode {
				t.Errorf("Expected status %v;  got %v", tc.expectedCode, res.Code)
			}

		})
	}
}
