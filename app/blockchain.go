package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

// Transaction within transaction value.
type Transaction struct {
	SenderBlockchainAddress string
	RecipientAddress        string
	Value                   float64
}

// Block is a mined some transactions.
type Block struct {
	PreviousHash string
	Timestamp    time.Time
	Nonce        int
	Transactions []Transaction
}

// Chain is a bucket to append mined block.
var Chain []Block
var transactionPool []Transaction

func init() {
	var initialHash []byte
	hash := sha256.Sum256(initialHash)
	CreateBlock(5, hex.EncodeToString(hash[:]))
}

// CreateBlock is create a struct based on args and transactions.
// And append created block to chain.
func CreateBlock(nonce int, ph string) {
	b := Block{
		PreviousHash: ph,
		Timestamp:    time.Now(),
		Nonce:        nonce,
		Transactions: transactionPool,
	}
	Chain = append(Chain, b)
}

// Hash is encrypt a block by sha256.
func (b *Block) Hash() string {
	bbyte, _ := json.Marshal(b)
	hash := sha256.Sum256(bbyte)
	return hex.EncodeToString(hash[:])
}

// AddTransaction is create a struct base on args.
// And append created transaction to transactionPool.
func AddTransaction(senderBlockchainAddress string, recipientAddress string, value float64) {
	ts := Transaction{
		SenderBlockchainAddress: senderBlockchainAddress,
		RecipientAddress:        recipientAddress,
		Value:                   value,
	}
	transactionPool = append(transactionPool, ts)
}
