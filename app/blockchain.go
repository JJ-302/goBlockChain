package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"regexp"
	"strings"
	"time"
)

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
const miningSender = "The BlockChain"
const miningReward = 1.0

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
func AddTransaction(senderAddress string, recipientAddress string, value float64) {
	tx := Transaction{
		SenderAddress:    senderAddress,
		RecipientAddress: recipientAddress,
		Value:            value,
	}
	transactionPool = append(transactionPool, tx)
}

// ValidProof is return whether was mining successfully.
func ValidProof(txs []Transaction, ph string, nonce int) bool {
	guessBlock := Block{
		PreviousHash: ph,
		Nonce:        nonce,
		Transactions: txs,
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

// Mining is run 'proof of work', create a block, and reward miner.
func Mining(blockchainAddress string) {
	AddTransaction(miningSender, blockchainAddress, miningReward)
	previousHash := Chain[len(Chain)-1].Hash()
	nonce := ProofOfWork()
	CreateBlock(nonce, previousHash)
}

// CalculateTotalAmount is calculates the amount of Bitcoin you have.
func CalculateTotalAmount(blockchainAddress string) float64 {
	totalAmount := 0.0
	for _, block := range Chain {
		for _, tx := range block.Transactions {
			if blockchainAddress == tx.RecipientAddress {
				totalAmount += tx.Value
			}
			if blockchainAddress == tx.SenderAddress {
				totalAmount -= tx.Value
			}
		}
	}
	return totalAmount
}
