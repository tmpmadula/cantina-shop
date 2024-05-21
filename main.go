package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tmpmadula/cantina-shop/db"
	"github.com/tmpmadula/cantina-shop/router"
)

func main() {
	// Connect to the database
	dbConn, err := db.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Create router
	r := router.NewRouter(dbConn)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(r)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
