package services

import (
	"encoding/hex"
	"fmt"

	"github.com/keybase/go-triplesec/sha3"
)

// GetEtheriumLikeAddressAsBytes method to get etherium like address.
func GetEtheriumLikeAddressAsBytes(srcData []byte) ([]byte, error) {
	if srcData == nil {
		return nil, fmt.Errorf("Data can't be null for address transformation")
	}

	digest := sha3.NewKeccak256()

	if _, err := digest.Write(srcData); err != nil {
		return nil, fmt.Errorf("Keccak256 can't digest a data %s", err)
	}

	return digest.Sum(nil)[12:], nil
}

// GetEtheriumLikeAddressFromCid ethereum address from cid
func GetEtheriumLikeAddressFromCid(cidData []byte) (string, error) {
	ethereumLikeAddress, err := GetEtheriumLikeAddressAsBytes(cidData)
	if err != nil {
		return "", fmt.Errorf("Can't hash id of cert %s", err.Error())
	}
	return hex.EncodeToString(ethereumLikeAddress), nil
}
