// internal/handlers/drink.go
package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tmpmadula/cantina-shop/internal/models"

	"github.com/gorilla/mux"
)

func GetDrinks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM drinks")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		drinks := []models.Drink{}
		for rows.Next() {
			var d models.Drink
			if err := rows.Scan(&d.ID, &d.Name, &d.Description, &d.Price, &d.Image, &d.Rating); err != nil {
				log.Fatal(err)
			}
			drinks = append(drinks, d)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(drinks)
	}
}

func GetDrink(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var d models.Drink
		err := db.QueryRow("SELECT * FROM drinks WHERE id = $1", id).Scan(&d.ID, &d.Name, &d.Description, &d.Price, &d.Image, &d.Rating)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(d)
	}
}

func CreateDrink(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d models.Drink
		json.NewDecoder(r.Body).Decode(&d)

		err := db.QueryRow("INSERT INTO drinks (name, description, price, image, rating) VALUES ($1, $2, $3, $4, 0) RETURNING id", d.Name, d.Description, d.Price, d.Image).Scan(&d.ID)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(d)
	}
}

func UpdateDrink(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d models.Drink
		json.NewDecoder(r.Body).Decode(&d)

		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("UPDATE drinks SET name = $1, description = $2, price = $3, image = $4, rating = $5 WHERE id = $6", d.Name, d.Description, d.Price, d.Image, d.Rating, id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(d)
	}
}

func DeleteDrink(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var d models.Drink
		err := db.QueryRow("SELECT * FROM drinks WHERE id = $1", id).Scan(&d.ID, &d.Name, &d.Description, &d.Price, &d.Image, &d.Rating)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = db.Exec("DELETE FROM drinks WHERE id = $1", id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode("Drink deleted")
	}
}
