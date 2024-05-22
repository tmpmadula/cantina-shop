package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tmpmadula/cantina-shop/internal/db"
	"github.com/tmpmadula/cantina-shop/internal/middleware"
	"github.com/tmpmadula/cantina-shop/router"
)

func main() {
	// Connect to database
	database, err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Initialize router
	r := router.NewRouter(database)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", middleware.JsonContentTypeMiddleware(r)))
}
