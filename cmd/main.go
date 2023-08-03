package main

import (
	"log"

	"github.com/AhmAlgiz/marketplace"
	"github.com/AhmAlgiz/marketplace/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	s := new(marketplace.Server)
	if err := s.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server running error: %s", err)
	}
}
