package blockchain

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"math/big"
)

type Transaction struct {
	From      []byte
	To        []byte
	Amount    float64
	Signature []byte
}

func (tx *Transaction) Hash() []byte {
	data := append(tx.From, tx.To...)
	data = append(data, []byte(fmt.Sprintf("%f", tx.Amount))...)

	hash := sha256.Sum256(data)
	return hash[:]
}

func (tx *Transaction) Sign(priv *ecdsa.PrivateKey) {
	hash := tx.Hash()

	r, s, err := ecdsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		return
	}

	signature := append(r.Bytes(), s.Bytes()...)
	tx.Signature = signature
}

func (tx *Transaction) Verify() bool {
	hash := tx.Hash()

	pubKeyInterface, err := x509.ParsePKIXPublicKey(tx.From)
	if err != nil {
		return false
	}

	pubKey, ok := pubKeyInterface.(*ecdsa.PublicKey)
	if !ok {
		return false
	}

	r := big.Int{}
	s := big.Int{}

	sigLen := len(tx.Signature)
	r.SetBytes(tx.Signature[:sigLen/2])
	s.SetBytes(tx.Signature[sigLen/2:])

	return ecdsa.Verify(pubKey, hash, &r, &s)
}
