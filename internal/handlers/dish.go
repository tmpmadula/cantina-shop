package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tmpmadula/cantina-shop/internal/models"

	"github.com/gorilla/mux"
)

func GetDishes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM dishes")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		dishes := []models.Dish{}
		for rows.Next() {
			var d models.Dish
			if err := rows.Scan(&d.ID, &d.Name, &d.Description, &d.Price, &d.Image, &d.Rating); err != nil {
				log.Fatal(err)
			}
			dishes = append(dishes, d)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(dishes)
	}
}

func GetDish(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var d models.Dish
		err := db.QueryRow("SELECT * FROM dishes WHERE id = $1", id).Scan(&d.ID, &d.Name, &d.Description, &d.Price, &d.Image, &d.Rating)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(d)
	}
}

func CreateDish(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d models.Dish
		json.NewDecoder(r.Body).Decode(&d)

		err := db.QueryRow("INSERT INTO dishes (name, description, price, image, rating) VALUES ($1, $2, $3, $4, 0) RETURNING id", d.Name, d.Description, d.Price, d.Image).Scan(&d.ID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(d)
	}
}

func UpdateDish(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d models.Dish
		json.NewDecoder(r.Body).Decode(&d)

		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("UPDATE dishes SET name = $1, description = $2, price = $3, image = $4, rating = $5 WHERE id = $6", d.Name, d.Description, d.Price, d.Image, d.Rating, id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(d)
	}
}

func DeleteDish(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var d models.Dish
		err := db.QueryRow("SELECT * FROM dishes WHERE id = $1", id).Scan(&d.ID, &d.Name, &d.Description, &d.Price, &d.Image, &d.Rating)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = db.Exec("DELETE FROM dishes WHERE id = $1", id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode("Dish deleted")
	}
}
