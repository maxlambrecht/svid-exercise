package validator

import (
	"testing"
	"crypto/tls"
	"crypto/x509"
	"path"
)

const assetsDir = "../test-assets"

func TestValidateID(t *testing.T) {
	x509Cert := helperLoadCertificate()

	testCases := []struct {
		name     string
		spiffeID string
		err      error
	}{
		{"Recognized Spiffe ID",
			"spiffe://example.com/service",
			nil,
			},
		{"Unrecognized Spiffe ID",
			"spiffe://example.com/other-service",
			ErrInvalidID},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			// Call the function to test
			err := SvidValidator{}.ValidateID(tc.spiffeID, x509Cert)

			// Check the Expectations
			if err != tc.err {
				t.Errorf("Expected %v;  got %v", tc.err, err)
			}

		})
	}

}

func helperLoadCertificate() *x509.Certificate {
	cert, _ := tls.LoadX509KeyPair(path.Join(assetsDir, "cert.pem"), path.Join(assetsDir, "key.pem"))
	certificate := []tls.Certificate{cert}
	x509Cert, _ := x509.ParseCertificate(certificate[0].Certificate[0])
	return x509Cert
}
