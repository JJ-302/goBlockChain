package main

import (
	"projects/goBlockChain/app"
	"projects/goBlockChain/config"
	"projects/goBlockChain/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	app.Printblock()
	app.CreateBlock(1, "second hash")
	app.Printblock()
	app.CreateBlock(2, "third hash")
	app.Printblock()
}
