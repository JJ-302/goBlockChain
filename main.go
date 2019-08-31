package main

import (
	"projects/goBlockChain/config"
	"projects/goBlockChain/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
}
