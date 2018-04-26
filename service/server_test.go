package service

import (
	"testing"
	"net/http/httptest"
	"path/filepath"
	"io/ioutil"
	"net/http"
	"bytes"
)

func TestHandler(t *testing.T) {

	cases := []struct {
		clientId string
		expectedCode int
	}{
		{"spiffe://dev.acme.com/path/service", 200},
		{"spiffe://dev.acme.com/path/unknown", 401},

	}

	for _, tc := range cases {
		cert := helperLoadCertificate(t)
		req, _ := http.NewRequest("GET", "/auth", bytes.NewReader(cert))

		res := httptest.NewRecorder()

		//Configure the service with the SPIFFE ID
		Config.SpiffeID = tc.clientId

		handle(res, req)

		if res.Code != tc.expectedCode{
			t.Errorf("expected status %v;  got %v", tc.expectedCode, res.Code)
		}

	}

}

func helperLoadCertificate(t *testing.T) []byte {
	path := filepath.Join("test-assets", "leaf.cert.pem")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
