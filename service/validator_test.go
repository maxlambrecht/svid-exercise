package service

import (
	"testing"
	"os"
)

func TestMatchCertificateID(t *testing.T) {

	f, err := os.Open("test-assets/leaf.cert.pem")
	if err != nil {
		t.Fatal("Could not open certificate")
	}
	defer f.Close()

	clientId := "spiffe://dev.acme.com/path/service"

	match, err := MatchCertificateID(clientId, f)

	if !match {
		t.Errorf("Expected match for %s did not match", clientId)
	}
}

