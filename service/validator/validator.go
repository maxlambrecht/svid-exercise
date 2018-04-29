package validator

import (
	"crypto/x509"
	"errors"
	"github.com/spiffe/go-spiffe/uri"
)


// ErrInvalidaID is returned when the Certificate does not contain
// a Spiffe ID that is valid
var ErrInvalidID = errors.New("Certificate ID is not valid ")

type Validator interface {
	ValidateID(id string, cert *x509.Certificate) error
}

type SvidValidator struct{}

// ValidateID validates that the certificate contains a Subject Alternative Name be equals to the id
// Returns an error if the Certificate cannot be parsed
// Returns ErrInvalidID if the id could not be found among the URI SANs
func (v SvidValidator) ValidateID(id string, cert *x509.Certificate) error {
	certIds, err := uri.GetURINamesFromCertificate(cert)
	if err != nil {
		return err
	}

	for _, certId := range certIds {
		if certId == id {
			return nil
		}
	}
	return ErrInvalidID
}
