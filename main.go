package main

import (
	"projects/goBlockChain/app"
	"projects/goBlockChain/config"
	"projects/goBlockChain/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	myBlockchainAddress := "myBlockchainAddress"

	app.AddTransaction("C", "D", 2.0)
	app.Mining(myBlockchainAddress)
	utils.Printblock()

	app.AddTransaction("E", "F", 3.0)
	app.AddTransaction("E", "G", 3.0)
	app.Mining(myBlockchainAddress)
	utils.Printblock()
}
