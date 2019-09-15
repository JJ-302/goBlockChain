package app

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
)

// Transaction is
type Transaction struct {
	SenderPrivateKey ecdsa.PrivateKey
	SenderPublicKey  ecdsa.PublicKey
	SenderAddress    string
	RecipientAddress string
	Value            float64
}

func (tx *Transaction) hash() []byte {
	txByte, _ := json.Marshal(tx)
	sha256Encoder := sha256.New()
	hash := sha256Encoder.Sum(txByte)
	return hash
}
