package main

import (
	"log"

	"github.com/SergioPaez96/twitter_react/bd"
	"github.com/SergioPaez96/twitter_react/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
