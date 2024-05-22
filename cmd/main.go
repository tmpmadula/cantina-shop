package main

import (
	"log"

	"github.com/tmpmadula/cantina-shop/internal/db"
	"github.com/tmpmadula/cantina-shop/internal/router"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Set up the router
	r := router.SetupRouter()

	// Run the server
	if err := r.Run(":8000"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
