package app

import (
	"crypto/sha256"
	"encoding/json"
)

// Transaction is bitcoin's transaction info.
type Transaction struct {
	SenderAddress    string
	RecipientAddress string
	Value            float64
}

func (tx *Transaction) hash() []byte {
	txByte, _ := json.Marshal(tx)
	sha256Encoder := sha256.New()
	sha256Encoder.Write(txByte)
	hash := sha256Encoder.Sum(nil)
	return hash
}

// AddTransaction returns result whether appending transaction successed.
func (tx *Transaction) AddTransaction(wallet *Wallet) bool {
	if tx.SenderAddress == MiningSender {
		TransactionPool = append(TransactionPool, *tx)
		return true
	}
	signature := wallet.GenerateSignature(tx.RecipientAddress, tx.Value)
	if signature.verifyTransactionSignature(&wallet.PublicKey, *tx) {
		TransactionPool = append(TransactionPool, *tx)
		return true
	}
	return false
}

// CreateTransaction returns new trunsaction struct.
func CreateTransaction(senAdd string, recAdd string, val float64) *Transaction {
	tx := Transaction{
		SenderAddress:    senAdd,
		RecipientAddress: recAdd,
		Value:            val,
	}
	return &tx
}
