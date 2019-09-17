package main

import (
	"fmt"
	"projects/goBlockChain/app"
	"projects/goBlockChain/config"
	"projects/goBlockChain/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	// myBlockchainAddress := "myBlockchainAddress"

	wallet1 := app.CreateWallet()
	wallet2 := app.CreateWallet()
	walletM := app.CreateWallet()

	tx := app.CreateTransaction(wallet1.BlockchainAddress, wallet2.BlockchainAddress, 1.5)
	result := tx.AddTransaction(wallet1)

	fmt.Println(result)

	app.Mining(walletM)
	utils.Printblock()

	tx = app.CreateTransaction(wallet1.BlockchainAddress, wallet2.BlockchainAddress, 4.5)
	result = tx.AddTransaction(wallet1)

	fmt.Println(result)
	app.Mining(walletM)
	utils.Printblock()

	fmt.Println("1", app.CalculateTotalAmount(wallet1.BlockchainAddress))
	fmt.Println("2", app.CalculateTotalAmount(wallet2.BlockchainAddress))
	fmt.Println("M", app.CalculateTotalAmount(walletM.BlockchainAddress))
}
