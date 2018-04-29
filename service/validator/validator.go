package validator

import (
	"io"
	"github.com/spiffe/go-spiffe/uri"
	"github.com/astaxie/beego/logs"
	"github.com/maxlambrecht/scytale-exercice/util"
)

// MatchCertificateID takes a string ID and X.509 certificate as a io.Reader
// and validates that the ID matches with an ID in the certificate
func MatchCertificateID(id string, cert io.Reader) (bool, error) {
	certIds, err := uri.FGetURINamesFromPEM(cert)
	if err != nil {
		logs.Error("Error at reading certificate ", err)
		return false, err
	}

	return util.Contains(certIds, id), nil
}

