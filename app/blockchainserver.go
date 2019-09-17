package app

import (
	"fmt"
	"log"
	"net/http"
	"projects/goDiary/config"
)

// StartBlockchainServer start blockchain node.
func StartBlockchainServer() error {
	log.Println("Port:8080 to listen on")
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
