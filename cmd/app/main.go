package main

import (
	"log"
	"myapp/internal/db"
	"myapp/internal/middleware"
	"myapp/router"
	"net/http"
	"os"
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
