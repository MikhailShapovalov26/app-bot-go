package main

import (
	config "alert/configs"
	app "alert/pkg/app"
	"flag"
	"log"
)

func main() {
	configPath := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	if *configPath == "" {
		log.Fatal("config path is required")
	}
	cfg, err := config.LoadFile(*configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration file: %v", err)
	}
	app.CreateServer(cfg)
}
