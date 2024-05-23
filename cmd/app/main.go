package main

import (
	"log"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/tmpmadula/cantina-shop/docs"
	"github.com/tmpmadula/cantina-shop/internal/db"
	"github.com/tmpmadula/cantina-shop/internal/middleware"
	"github.com/tmpmadula/cantina-shop/router"
)

// @title Cantina Shop API
// @version 1.0
// @description This is a sample server for a cantina shop.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api

func main() {
	// Connect to database
	database, err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Initialize router
	r := router.NewRouter(database)

	// Serve Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", middleware.JsonContentTypeMiddleware(r)))
}
