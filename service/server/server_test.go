package server

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
	"github.com/maxlambrecht/svid-exercise/service/validator"
	"log"
)

const assetsDir = "../test-assets"

func TestAuthenticateHandler(t *testing.T) {

	authServer := AuthServer{
		CertValidator: validator.SvidValidator{},
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

// Integration Test for testing that the HTTPS server handle and validate correctly
// the SpiffeID in the SVID x509 certificates sent by the clients using TLS
func TestHTTPSServer(t *testing.T) {

	cert := path.Join(assetsDir, "cert.pem")
	key := path.Join(assetsDir, "key.pem")
	untrustedCert := path.Join(assetsDir, "untrusted_cert.pem")
	untrustedKey := path.Join(assetsDir, "untrusted_key.pem")

	authServer := AuthServer{
		Addr:          ":3457",
		SpiffeID:      "spiffe://example.com/service",
		CertFile:      cert,
		KeyFile:       key,
		CertValidator: validator.SvidValidator{},
	}

	// Run the HTTPS server
	go func() {
		authServer.Start()
	}()

	// Create a client with a trusted SpiffeID
	client := createClient(cert, key)

	// Perform a the authentication request
	res, err := client.Get("https://localhost:3457/auth")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Response code was %v; want 200", res.StatusCode)
	}



	// Create a client with an untrusted SpiffeID
	client = createClient(untrustedCert, untrustedKey)

	// Perform a the authentication request
	res2, err := client.Get("https://localhost:3457/auth")
	if err != nil {
		t.Fatal(err)
	}
	defer res2.Body.Close()

	if res2.StatusCode != 401 {
		t.Errorf("Response code was %v; want 401", res.StatusCode)
	}

	// Send signal to shutdown the server
	shutdown<-1
}

func createClient(cert string, key string) *http.Client {
	x509, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{x509},
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	return &http.Client{Transport: transport}
}

func helperLoadCertificate() *x509.Certificate {
	cert, _ := tls.LoadX509KeyPair(path.Join(assetsDir, "cert.pem"), path.Join(assetsDir, "key.pem"))
	certificate := []tls.Certificate{cert}
	x509Cert, _ := x509.ParseCertificate(certificate[0].Certificate[0])
	return x509Cert
}
