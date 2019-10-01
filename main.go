package main

import (
	"flag"
	"fmt"
	"log"
	"projects/goBlockChain/app"
	"projects/goBlockChain/config"
	"projects/goBlockChain/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)

	port := flag.Int("int", 8080, "port")
	flag.Parse()

	addr := utils.GetHost()
	fmt.Println(utils.FindNeighbours(addr, *port, 0, 2, 8080, 8082))

	minerWallet := app.CreateWallet()
	go app.StartMining(minerWallet)

	log.Println(app.StartBlockchainServer(*port))
}
