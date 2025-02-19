package main

import (
	"fmt"
	"log"

	"github.com/erobx/tradeups/backend/internal/db"
	"github.com/erobx/tradeups/backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := db.NewPostgresDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server...")
	s := server.NewServer("8080", db)

	if err := s.Run(); err != nil {
		log.Fatalf("Server could not be ran: %v", err)
	}
}


