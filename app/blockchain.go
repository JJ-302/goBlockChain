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

func Printblock() {
	headerLine := strings.Repeat("=", 25)
	for i, v := range chain {
	format := "%-15s : %v\n"
		fmt.Println(headerLine + "Chain" + strconv.Itoa(i) + headerLine)
		fmt.Printf(format, "PreviousHash", v.PreviousHash)
		fmt.Printf(format, "Timestamp", v.Timestamp.Format(time.RFC3339))
		fmt.Printf(format, "Nonce", v.Nonce)
		fmt.Printf(format, "Transactions", v.Transactions)
	}
	fmt.Println(strings.Repeat("*", 50))
}
