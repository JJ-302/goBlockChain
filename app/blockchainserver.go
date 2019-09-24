package app

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"projects/goBlockChain/config"
)

func init() {
	var initialHash []byte
	hash := sha256.Sum256(initialHash)
	CreateBlock(5, hex.EncodeToString(hash[:]), TransactionPool)
}

// StartBlockchainServer start blockchain node.
func StartBlockchainServer() error {
	log.Println("Port:8080 to listen on")
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
