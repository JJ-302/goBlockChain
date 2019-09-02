package main

import (
	"projects/goBlockChain/app"
	"projects/goBlockChain/config"
	"projects/goBlockChain/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	utils.Printblock()

	app.AddTransaction("A", "B", 1.0)
	previousHash := app.Chain[len(app.Chain)-1].Hash()
	app.CreateBlock(1, previousHash)
	utils.Printblock()

	app.AddTransaction("C", "D", 2.0)
	app.AddTransaction("E", "F", 3.0)
	previousHash = app.Chain[len(app.Chain)-1].Hash()
	app.CreateBlock(2, previousHash)
	utils.Printblock()
}
