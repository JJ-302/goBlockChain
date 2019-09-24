package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
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

// TransactionPool is
var TransactionPool []Transaction

const miningDifficulty = 3

// MiningSender is send reward to miner.
const MiningSender = "The BlockChain"
const miningReward = 1.0

// CreateBlock is create a struct based on args and transactions.
// And append created block to chain.
func CreateBlock(nonce int, ph string, txs []Transaction) {
	b := Block{
		PreviousHash: ph,
		Timestamp:    time.Now(),
		Nonce:        nonce,
		Transactions: txs,
	}
	Chain = append(Chain, b)
	TransactionPool = TransactionPool[:0]
}

func (b *Block) hash() string {
	bbyte, _ := json.Marshal(b)
	hash := sha256.Sum256(bbyte)
	return hex.EncodeToString(hash[:])
}

func validProof(txs []Transaction, ph string, nonce int) bool {
	guessBlock := Block{
		PreviousHash: ph,
		Nonce:        nonce,
		Transactions: txs,
	}
	guessHash := guessBlock.hash()
	validHash := regexp.MustCompile("^" + strings.Repeat("0", miningDifficulty))
	return validHash.MatchString(guessHash)
}

func proofOfWork() (int, []Transaction) {
	transactions := make([]Transaction, len(TransactionPool))
	copy(transactions, TransactionPool)
	previousHash := Chain[len(Chain)-1].hash()
	nonce := 0
	for !validProof(transactions, previousHash, nonce) {
		nonce++
	}
	return nonce, transactions
}

// Mining is run 'proof of work', create a block, and reward miner.
func Mining(wallet *Wallet) {
	tx := CreateTransaction(MiningSender, wallet.BlockchainAddress, miningReward)
	if !tx.AddTransaction(wallet) {
		log.Fatalln("exit!")
	}
	previousHash := Chain[len(Chain)-1].hash()
	nonce, txs := proofOfWork()
	CreateBlock(nonce, previousHash, txs)
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
