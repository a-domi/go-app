package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Logging string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Faild to read ini file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Logging: cfg.Section("app").Key("logging").String(),
	}
}
