package main

import (
	"log"

	"github.com/moomerman/go-shorty/shorty"
)

func main() {
	if err := shorty.LoadConfig(); err != nil {
		log.Fatal("unable to load config", err)
	}

	log.Fatal(shorty.Start())
}
