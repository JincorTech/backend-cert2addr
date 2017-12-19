package services

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// GetCIDFromPem get cid from cert. in pem format
func GetCIDFromPem(pemData []byte) (string, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return "", errors.New("Expecting a PEM-encoded X509 certificate; PEM block not found")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", errors.New("failed to parse certificate")
	}

	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("x509::%s::%s", getDN(&cert.Subject), getDN(&cert.Issuer)))), nil
}
