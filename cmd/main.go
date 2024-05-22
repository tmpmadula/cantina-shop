// cmd/main.go
package main

import (
	"github.com/tmpmadula/cantina-shop/internal/db"
	"github.com/tmpmadula/cantina-shop/internal/router"
)

func main() {
	// Initialize database
	db.InitDB()

	// Set up router
	r := router.SetupRouter()

	// Run the server
	r.Run(":8080")
}
