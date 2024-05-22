// router/router.go
package router

import (
	"database/sql"

	"github.com/tmpmadula/cantina-shop/internal/handlers"
	"github.com/tmpmadula/cantina-shop/internal/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", handlers.RegisterUser(db)).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser(db)).Methods("POST")

	// Protected routes
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)
	api.Use(middleware.RateLimit)
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
