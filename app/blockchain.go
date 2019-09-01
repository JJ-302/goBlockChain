package app

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Value float64
}

type Block struct {
	PreviousHash string
	Timestamp    time.Time
	Nonce        int
	Transactions []Transaction
}

var chain []Block
var transactionPool []Transaction

func init() {
	CreateBlock(5, "initial hash")
}

func CreateBlock(nonce int, ph string) {
	b := Block{
		PreviousHash: ph,
		Timestamp:    time.Now(),
		Nonce:        nonce,
		Transactions: transactionPool,
	}
	chain = append(chain, b)
}

