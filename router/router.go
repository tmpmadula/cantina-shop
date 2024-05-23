// router/router.go
package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/tmpmadula/cantina-shop/internal/auth"
	"github.com/tmpmadula/cantina-shop/internal/handlers"
	"github.com/tmpmadula/cantina-shop/internal/middleware"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", handlers.RegisterUser(db)).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser(db)).Methods("POST")

	router.HandleFunc("/auth/google/login", auth.HandleGoogleLogin).Methods("GET")
	router.HandleFunc("/auth/google/callback", auth.HandleGoogleCallback(db)).Methods("GET")

	// Protected routes
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.Use(middleware.RateLimitMiddleware)
	router.HandleFunc("/users", handlers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser(db)).Methods("DELETE")

	router.HandleFunc("/dishes", handlers.GetDishes(db)).Methods("GET")
	router.HandleFunc("/dishes/{id}", handlers.GetDish(db)).Methods("GET")
	router.HandleFunc("/dishes", handlers.CreateDish(db)).Methods("POST")
	router.HandleFunc("/dishes/{id}", handlers.UpdateDish(db)).Methods("PUT")
	router.HandleFunc("/dishes/{id}", handlers.DeleteDish(db)).Methods("DELETE")

	router.HandleFunc("/drinks", handlers.GetDrinks(db)).Methods("GET")
	router.HandleFunc("/drinks/{id}", handlers.GetDrink(db)).Methods("GET")
	router.HandleFunc("/drinks", handlers.CreateDrink(db)).Methods("POST")
	router.HandleFunc("/drinks/{id}", handlers.UpdateDrink(db)).Methods("PUT")
	router.HandleFunc("/drinks/{id}", handlers.DeleteDrink(db)).Methods("DELETE")

	return router
}

/*

// router/router.go
package router

import (
    "database/sql"
    "myapp/internal/handlers"
    "myapp/internal/middleware"
    "github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/register", handlers.RegisterUser(db)).Methods("POST")
    router.HandleFunc("/login", handlers.LoginUser(db)).Methods("POST")

    api := router.PathPrefix("/api").Subrouter()
    api.Use(middleware.AuthMiddleware)
    api.Use(middleware.RateLimitMiddleware)

    api.HandleFunc("/users", handlers.GetUsers(db)).Methods("GET")
    api.HandleFunc("/users/{id}", handlers.GetUser(db)).Methods("GET")
    api.HandleFunc("/users", handlers.CreateUser(db)).Methods("POST")
    api.HandleFunc("/users/{id}", handlers.UpdateUser(db)).Methods("PUT")
    api.HandleFunc("/users/{id}", handlers.DeleteUser(db)).Methods("DELETE")

    api.HandleFunc("/dishes", handlers.GetDishes(db)).Methods("GET")
    api.HandleFunc("/dishes/{id}", handlers.GetDish(db)).Methods("GET")
    api.HandleFunc("/dishes", handlers.CreateDish(db)).Methods("POST")
    api.HandleFunc("/dishes/{id}", handlers.UpdateDish(db)).Methods("PUT")
    api.HandleFunc("/dishes/{id}", handlers.DeleteDish(db)).Methods("DELETE")

    api.HandleFunc("/drinks", handlers.GetDrinks(db)).Methods("GET")
    api.HandleFunc("/drinks/{id}", handlers.GetDrink(db)).Methods("GET")
    api.HandleFunc("/drinks", handlers.CreateDrink(db)).Methods("POST")
    api.HandleFunc("/drinks/{id}", handlers.UpdateDrink(db)).Methods("PUT")
    api.HandleFunc("/drinks/{id}", handlers.DeleteDrink(db)).Methods("DELETE")

    api.HandleFunc("/reviews", handlers.GetReviews(db)).Methods("GET")
    api.HandleFunc("/reviews/{id}", handlers.GetReview(db)).Methods("GET")
    api.HandleFunc("/reviews", handlers.CreateReview(db)).Methods("POST")
    api.HandleFunc("/reviews/{id}", handlers.UpdateReview(db)).Methods("PUT")
    api.HandleFunc("/reviews/{id}", handlers.DeleteReview(db)).Methods("DELETE")

    return router
}

*/
