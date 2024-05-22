package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/tmpmadula/cantina-shop/handler"
)

func NewRouter(db *sql.DB) *mux.Router {
	userHandler := &handler.UserHandler{DB: db}
	dishHandler := &handler.DishHandler{DB: db}
	drinkHandler := &handler.DrinkHandler{DB: db}
	reviewHandler := &handler.ReviewHandler{DB: db}
	authHandler := &handler.AuthHandler{DB: db}

	router := mux.NewRouter()

	// User endpoints
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Dish endpoints
	router.HandleFunc("/dishes", dishHandler.GetDishes).Methods("GET")
	router.HandleFunc("/dishes/{id}", dishHandler.GetDish).Methods("GET")
	router.HandleFunc("/dishes", dishHandler.CreateDish).Methods("POST")
	router.HandleFunc("/dishes/{id}", dishHandler.UpdateDish).Methods("PUT")
	router.HandleFunc("/dishes/{id}", dishHandler.DeleteDish).Methods("DELETE")

	// Drink endpoints
	router.HandleFunc("/drinks", drinkHandler.GetDrinks).Methods("GET")
	router.HandleFunc("/drinks/{id}", drinkHandler.GetDrink).Methods("GET")
	router.HandleFunc("/drinks", drinkHandler.CreateDrink).Methods("POST")
	router.HandleFunc("/drinks/{id}", drinkHandler.UpdateDrink).Methods("PUT")
	router.HandleFunc("/drinks/{id}", drinkHandler.DeleteDrink).Methods("DELETE")

	// Review endpoints
	router.HandleFunc("/reviews", reviewHandler.GetReviews).Methods("GET")
	router.HandleFunc("/reviews/{id}", reviewHandler.GetReview).Methods("GET")
	router.HandleFunc("/reviews", reviewHandler.CreateReview).Methods("POST")
	router.HandleFunc("/reviews/{id}", reviewHandler.UpdateReview).Methods("PUT")
	router.HandleFunc("/reviews/{id}", reviewHandler.DeleteReview).Methods("DELETE")

	// Authentication endpoints
	router.HandleFunc("/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	// OAuth2 integration endpoints
	router.HandleFunc("/oauth/google", authHandler.GoogleOAuthCallback).Methods("GET")

	return router
}
