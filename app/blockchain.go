package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"regexp"
	"strings"
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

const miningDifficulty = 3

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

// ValidProof is return whether was mining successfully.
func ValidProof(transactions []Transaction, ph string, nonce int) bool {
	guessBlock := Block{
		PreviousHash: ph,
		Nonce:        nonce,
		Transactions: transactions,
	}
	guessHash := guessBlock.Hash()
	validHash := regexp.MustCompile("^" + strings.Repeat("0", miningDifficulty))
	return validHash.MatchString(guessHash)
}

// ProofOfWork is return nonce when success mining.
func ProofOfWork() int {
	transactions := make([]Transaction, len(transactionPool))
	copy(transactions, transactionPool)
	previousHash := Chain[len(Chain)-1].Hash()
	nonce := 0
	for !ValidProof(transactions, previousHash, nonce) {
		nonce++
	}
	return nonce
}
