package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
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

// Printblock is print blockchain on terminal.
func Printblock() {
	headerLine := strings.Repeat("=", 25)
	formatA := "%-15s : %v\n"
	formatB := "%-25s : %v\n"
	for i, v := range Chain {
		fmt.Println(headerLine + "Chain" + strconv.Itoa(i) + headerLine)
		fmt.Printf(formatA, "PreviousHash", v.PreviousHash)
		fmt.Printf(formatA, "Timestamp", v.Timestamp.Format(time.RFC3339))
		fmt.Printf(formatA, "Nonce", v.Nonce)
		fmt.Println("Transactions")
		fmt.Println(strings.Repeat("-", 50))
		for _, v := range v.Transactions {
			fmt.Printf(formatB, "SenderBlockchainAddress", v.SenderBlockchainAddress)
			fmt.Printf(formatB, "RecipientAddress", v.RecipientAddress)
			fmt.Printf(formatB, "Value", v.Value)
		}
	}
	fmt.Printf("%s\n\n\n", strings.Repeat("*", 50))
}
