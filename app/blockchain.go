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
	Value float64
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

// Printblock is print blockchain on terminal.
func Printblock() {
	headerLine := strings.Repeat("=", 25)
	format := "%-15s : %v\n"
	for i, v := range Chain {
		fmt.Println(headerLine + "Chain" + strconv.Itoa(i) + headerLine)
		fmt.Printf(format, "PreviousHash", v.PreviousHash)
		fmt.Printf(format, "Timestamp", v.Timestamp.Format(time.RFC3339))
		fmt.Printf(format, "Nonce", v.Nonce)
		fmt.Printf(format, "Transactions", v.Transactions)
	}
	fmt.Println(strings.Repeat("*", 50))
}
