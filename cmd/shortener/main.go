package main

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("error running server: ", err)
	}

	if err = app.Run(cfg); err != nil {
		log.Fatal("error running server", err)
	}
}
