package main

import (
	"projects/goBlockChain/app"
	"projects/goBlockChain/config"
	"projects/goBlockChain/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	app.Printblock()

	previousHash := app.Chain[len(app.Chain)-1].Hash()
	app.CreateBlock(1, previousHash)
	app.Printblock()

	previousHash = app.Chain[len(app.Chain)-1].Hash()
	app.CreateBlock(2, previousHash)
	app.Printblock()
}
