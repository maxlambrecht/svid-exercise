package validator

import (
	"crypto/x509"
	"errors"
	"github.com/spiffe/go-spiffe/uri"
)

type Validator interface {
	ValidateID(id string, cert *x509.Certificate) error
}

// ErrInvalidaID is returned when the Certificate does not contain
// a Spiffe ID that is recognized by the server
var ErrInvalidID = errors.New("Certificate ID is not valid ")

// SvidValidator implements the Validator interface to validate
// SVID certificates
type SvidValidator struct{}

// ValidateID validates that the certificate contains a Subject Alternative Name be equals to the spiffeID
// Returns an error if the Certificate cannot be parsed
// Returns ErrInvalidID if the id could not be found among the URI SANs
func (v SvidValidator) ValidateID(spiffeID string, cert *x509.Certificate) error {
	certIds, err := uri.GetURINamesFromCertificate(cert)
	if err != nil {
		return err
	}

	for _, certId := range certIds {
		if certId == spiffeID {
			return nil
		}
	}
	return ErrInvalidID
}
