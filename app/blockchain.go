package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"projects/goBlockChain/utils"
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

// Neighbours is neighbour node Address.
var Neighbours []string

const miningDifficulty = 3
const portRangeStart = 8080
const portRangeEnd = 8082
const ipRangeStart = 0
const ipRangeEnd = 1
const neighboursSyncTimeSec = 20

// MiningSender is send reward to miner.
const MiningSender = "The BlockChain"
const miningReward = 1.0

// SetNeighbours is
func SetNeighbours(port int) {
	addr := utils.GetHost()
	for {
		Neighbours = utils.FindNeighbours(addr, port, ipRangeStart, ipRangeEnd, portRangeStart, portRangeEnd)
		log.Println("neighbours: ", Neighbours)
		time.Sleep(neighboursSyncTimeSec * time.Second)
	}
}

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
	for _, node := range Neighbours {
		url := "http://" + node + "/sync/delete"
		resp, err := http.Post(url, "application/json", nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()
	}
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
	if len(TransactionPool) == 0 {
		fmt.Println("Transaction is empty.")
		return
	}
	tx := CreateTransaction(MiningSender, wallet.BlockchainAddress, miningReward)
	if !tx.AddTransaction(wallet) {
		log.Fatalln("exit!")
	}
	previousHash := Chain[len(Chain)-1].hash()
	nonce, txs := proofOfWork()
	CreateBlock(nonce, previousHash, txs)
	for _, node := range Neighbours {
		url := "http://" + node + "/consensus"
		resp, err := http.Post(url, "application/json", nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()
	}
	fmt.Println("Mining is successfully")
}

func validChain(chain []Block) bool {
	preBlock := chain[0]
	currentIndex := 1
	for i := currentIndex; i < len(chain); i++ {
		block := chain[currentIndex]
		if block.PreviousHash != preBlock.hash() {
			return false
		}
		if !validProof(block.Transactions, block.PreviousHash, block.Nonce) {
			return false
		}
		preBlock = block
		currentIndex++
	}
	return true
}

func ResolveConflicts() bool {
	maxLength := len(Chain)
	for _, node := range Neighbours {
		url := "http://" + node + "/chain"
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
			return false
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return false
		}
		var respChain []Block
		if err = json.Unmarshal(body, &respChain); err != nil {
			log.Println(err)
			return false
		}
		fmt.Println(respChain)
		chainLength := len(respChain)
		if chainLength > maxLength && validChain(respChain) {
			maxLength = chainLength
			Chain = respChain
		}
	}
	if len(Chain) == 0 {
		log.Println("chain_was_not_replaced")
		return false
	}
	log.Println("chain_was_replaced")
	return true
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
