package main

import (
	"log"
	"net/http"
	"os"

	"github.com/juju/ratelimit"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/tmpmadula/cantina-shop/docs"
	"github.com/tmpmadula/cantina-shop/internal/db"
	"github.com/tmpmadula/cantina-shop/internal/middleware"
	"github.com/tmpmadula/cantina-shop/router"
	"github.com/urfave/negroni"
)

func rateLimitMiddleware(bucket *ratelimit.Bucket) negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if bucket.TakeAvailable(1) == 0 {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next(w, r)
	})
}

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

	// Use Negroni for middleware
	n := negroni.Classic()

	// Add rate limiting middleware
	bucket := ratelimit.NewBucketWithRate(10, 10) // 10 requests per second
	n.Use(rateLimitMiddleware(bucket))

	// Initialize router
	r := router.NewRouter(database)

	// Serve Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":8000", middleware.JsonContentTypeMiddleware(r)))
}
