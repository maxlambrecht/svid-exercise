package server

import (
	"testing"
	"net/http/httptest"
	"crypto/tls"
	"crypto/x509"
	"path"
	"net/http"
)

const assetsDir = "../test-assets"

func TestAuthenticateHandler(t *testing.T) {

	authServer := AuthServer{
		Validator: CertValidator{},
	}

	// Create the client request that will be used in the test cases and inject the SVID
	request, _ := http.NewRequest("GET", "/auth", nil)
	x509Cert := helperLoadCertificate()
	request.TLS = &tls.ConnectionState{}
	request.TLS.PeerCertificates = make([]*x509.Certificate, 1)
	request.TLS.PeerCertificates[0] = x509Cert


	testCases := []struct {
		name         string
		spiffeID     string
		expectedCode int
	}{
		{"valid ID", "spiffe://example.com/service", 200},
		{"invalid ID", "spiffe://example.com/other-service", 401},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			// Configure the service with the SPIFFE ID
			authServer.SpiffeID = tc.spiffeID

			res := httptest.NewRecorder()

			// Call the function to test
			authServer.authenticateHandler(res, request)

			// Check the Expectations
			if res.Code != tc.expectedCode {
				t.Errorf("Expected status %v;  got %v", tc.expectedCode, res.Code)
			}

		})
	}
}

func helperLoadCertificate() (*x509.Certificate) {
	cert, _ := tls.LoadX509KeyPair(path.Join(assetsDir, "cert.pem"), path.Join(assetsDir, "key.pem"))
	certificate := []tls.Certificate{cert}
	x509Cert, _ := x509.ParseCertificate(certificate[0].Certificate[0])
	return x509Cert
}
