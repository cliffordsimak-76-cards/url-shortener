package main

import (
	"log"
	_ "net/http/pprof"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	printBuildData()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("error running server: ", err)
	}

	if err = app.Run(cfg); err != nil {
		log.Fatal("error running server", err)
	}
}

func printBuildData() {
	log.Printf("Build version: %v\n", buildVersion)
	log.Printf("Build date: %v\n", buildDate)
	log.Printf("Build commit: %v\n", buildCommit)
}
