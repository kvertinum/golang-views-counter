package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Kvertinum01/views-counter/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	conf := &server.Config{}
	if _, err := toml.DecodeFile(configPath, conf); err != nil {
		log.Fatal(err)
	}

	if err := server.SetupServer(conf); err != nil {
		log.Fatal(err)
	}
}
