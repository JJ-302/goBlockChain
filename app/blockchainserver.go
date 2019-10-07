package app

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

const methodPost = "POST"

func init() {
	var initialHash []byte
	hash := sha256.Sum256(initialHash)
	CreateBlock(5, hex.EncodeToString(hash[:]), TransactionPool)
}

func getChainHandler(w http.ResponseWriter, r *http.Request) {
	jsonValue, _ := json.Marshal(Chain)
	w.Write(jsonValue)
}

func createWalletHandler(w http.ResponseWriter, r *http.Request) {
	toAllowAccess(w)
	jsonValue, _ := json.Marshal(CreateWallet())
	w.Write(jsonValue)
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates := template.Must(template.ParseFiles("app/views/transaction.html"))

		if err := templates.Execute(w, TransactionPool); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	toAllowAccess(w)

	if r.Method == methodPost {
		var tx Transaction
		body := parseJSON(r)
		if err := json.Unmarshal(body, &tx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, isExist := WalletPool[tx.RecipientAddress]; !isExist {
			writeResponse(w, false)
			return
		}

		if wallet := WalletPool[tx.SenderAddress]; tx.addTransaction(&wallet) {
			for _, node := range Neighbours {
				url := "http://" + node + "/sync"
				resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
				if err != nil {
					log.Println(err)
				}
				defer resp.Body.Close()
			}
			writeResponse(w, true)
		} else {
			writeResponse(w, false)
		}
	}
}

func syncTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == methodPost {
		var tx Transaction
		body := parseJSON(r)
		if err := json.Unmarshal(body, &tx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tx.syncTransaction()
	}
}

func deleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == methodPost {
		TransactionPool = TransactionPool[:0]
	}
}

func consensusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == methodPost {
		result := ResolveConflicts()
		log.Println("Resolve conflicts:", result)
	}
}

func calcTotalAmountHandler(w http.ResponseWriter, r *http.Request) {
	toAllowAccess(w)

	if r.Method == methodPost {
		var wallet Wallet
		body := parseJSON(r)
		if err := json.Unmarshal(body, &wallet); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		amount := CalculateTotalAmount(wallet.BlockchainAddress)
		jsonValue, _ := json.Marshal(map[string]float64{"result": amount})
		w.Write(jsonValue)
	}
}

func toAllowAccess(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

func writeResponse(w http.ResponseWriter, result bool) {
	jsonValue, _ := json.Marshal(map[string]bool{"result": result})
	w.Write(jsonValue)
}

func parseJSON(r *http.Request) []byte {
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		log.Fatalln(err)
	}
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil {
		log.Println(err)
	}
	return body[:length]
}

// StartMining is mining every 10sec.
func StartMining(wallet *Wallet) {
	fmt.Println("mining to listen on")
	for {
		Mining(wallet)
		time.Sleep(10000 * time.Millisecond)
	}
}

// StartBlockchainServer start blockchain node.
func StartBlockchainServer(port int) error {
	log.Printf("Port: %v to listen on", port)
	http.HandleFunc("/wallet", createWalletHandler)
	http.HandleFunc("/chain", getChainHandler)
	http.HandleFunc("/transaction", transactionHandler)
	http.HandleFunc("/sync", syncTransactionHandler)
	http.HandleFunc("/sync/delete", deleteTransactionHandler)
	http.HandleFunc("/consensus", consensusHandler)
	http.HandleFunc("/calc", calcTotalAmountHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
