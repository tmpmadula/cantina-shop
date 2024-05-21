package router

import (
	"database/sql"

	"github.com/gorilla/mux"

	"github.com/tmpmadula/cantina-shop/handler"
)

func NewRouter(db *sql.DB) *mux.Router {
	userHandler := &handler.UserHandler{DB: db}

	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	return router
}
