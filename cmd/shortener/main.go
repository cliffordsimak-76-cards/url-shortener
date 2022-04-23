package main

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal("error running server ", err)
	}
}
