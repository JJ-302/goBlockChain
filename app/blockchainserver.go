package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"projects/goBlockChain/config"
	"strconv"
)

func init() {
	var initialHash []byte
	hash := sha256.Sum256(initialHash)
	CreateBlock(5, hex.EncodeToString(hash[:]), TransactionPool)
}

func getChainHandler(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("app/views/chain.html"))
	err := templates.Execute(w, Chain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createWalletHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	wallet := CreateWallet()
	jsonValue, _ := json.Marshal(wallet)
	w.Write(jsonValue)
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates := template.Must(template.ParseFiles("app/views/transaction.html"))
		err := templates.Execute(w, TransactionPool)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	if r.Method == "POST" {
		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		body := make([]byte, length)
		length, err = r.Body.Read(body)

		var tx Transaction
		err = json.Unmarshal(body[:length], &tx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		wallet := WalletPool[tx.SenderAddress]
		if tx.AddTransaction(&wallet) {
			jsonValue, _ := json.Marshal(map[string]bool{"result": true})
			w.Write(jsonValue)
		} else {
			jsonValue, _ := json.Marshal(map[string]bool{"result": false})
			w.Write(jsonValue)
		}
	}
}

// StartBlockchainServer start blockchain node.
func StartBlockchainServer() error {
	log.Println("Port:8080 to listen on")
	http.HandleFunc("/wallet", createWalletHandler)
	http.HandleFunc("/chain", getChainHandler)
	http.HandleFunc("/transaction", transactionHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
