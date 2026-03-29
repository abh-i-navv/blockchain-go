package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
)

type Wallet struct {
	Private *ecdsa.PrivateKey
	Public  []byte
}

func NewWallet() *Wallet {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	pubBytes, _ := x509.MarshalPKIXPublicKey(&priv.PublicKey)

	return &Wallet{
		Private: priv,
		Public:  pubBytes,
	}
}
