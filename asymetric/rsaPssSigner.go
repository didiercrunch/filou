package asymetric

import (
	_ "crypto/sha256"
	_ "crypto/sha512"

	"crypto"
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
)

type RsaPssSigner struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	Hash       crypto.Hash

	Random io.Reader
}

func (this *RsaPssSigner) getHash() (crypto.Hash, error) {
	if !this.Hash.Available() {
		return 0, errors.New(fmt.Sprintf("hash function %v is unabailable", this.Hash))
	} else {
		return this.Hash, nil
	}
}

func (this *RsaPssSigner) Sign(hashedData []byte) ([]byte, error) {
	ret, err := rsa.SignPSS(this.Random, this.PrivateKey, this.Hash, hashedData, nil)
	return ret, err
}

func (this *RsaPssSigner) VerifySignature(hashedData []byte, signature []byte) error {
	return rsa.VerifyPSS(this.PublicKey, this.Hash, hashedData, signature, nil)
}
