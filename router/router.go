package router

import (
	"database/sql"

	"github.com/tmpmadula/cantina-shop/handler"

	"github.com/gorilla/mux"
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
