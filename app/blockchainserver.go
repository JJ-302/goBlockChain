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
)

var templates = template.Must(template.ParseFiles("app/views/chain.html"))

func init() {
	var initialHash []byte
	hash := sha256.Sum256(initialHash)
	CreateBlock(5, hex.EncodeToString(hash[:]), TransactionPool)
}

func getChainHandler(w http.ResponseWriter, r *http.Request) {
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

// StartBlockchainServer start blockchain node.
func StartBlockchainServer() error {
	log.Println("Port:8080 to listen on")
	http.HandleFunc("/wallet", createWalletHandler)
	http.HandleFunc("/chain", getChainHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
