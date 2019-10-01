package main

import (
	"log"
	"projects/goBlockChain/app"
	"projects/goBlockChain/config"
	"projects/goBlockChain/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	minerWallet := app.CreateWallet()
	go app.StartMining(minerWallet)
	log.Println(app.StartBlockchainServer())
}
