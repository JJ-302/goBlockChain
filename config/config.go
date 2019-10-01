package config

import (
	"log"
	"os"

	ini "gopkg.in/ini.v1"
)

// Configlist has some info for application.
type Configlist struct {
	Port    int
	LogFile string
}

// Config is result when initialized Configlist.
var Config Configlist

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = Configlist{
		Port:    cfg.Section("web").Key("port").MustInt(),
		LogFile: cfg.Section("goblockchain").Key("log_file").String(),
	}
}
