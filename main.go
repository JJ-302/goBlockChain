package main

import (
	"flag"
	"log"
	"projects/goBlockChain/app"
	"projects/goBlockChain/config"
	"projects/goBlockChain/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)

	port := flag.Int("int", 8080, "port")
	flag.Parse()

	minerWallet := app.CreateWallet()
	go app.StartMining(minerWallet)

	log.Println(app.StartBlockchainServer(*port))
}
