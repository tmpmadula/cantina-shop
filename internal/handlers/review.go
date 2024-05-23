package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/tmpmadula/cantina-shop/internal/models"
)

func GetReviews(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, user_id, dish_id, drink_id, rating, comment FROM reviews")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var reviews []models.Review
		for rows.Next() {
			var review models.Review
			if err := rows.Scan(&review.ID, &review.UserID, &review.DrinkID, &review.DishID, &review.Rating, &review.Comment); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			reviews = append(reviews, review)
		}

		json.NewEncoder(w).Encode(reviews)
	}

}

// GetReview

func GetReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		var review models.Review
		err := db.QueryRow("SELECT id, user_id, dish_id, drink_id, rating, comment FROM reviews WHERE id = $1", id).
			Scan(&review.ID, &review.UserID, &review.DishID, &review.DrinkID, &review.Rating, &review.Comment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(review)
	}
}

// CreateReview

func CreateReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var review models.Review
		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if review.UserID == 0 || (review.DishID == 0 && review.DrinkID == 0) || review.Rating == 0 {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		var id int
		err := db.QueryRow("INSERT INTO reviews (user_id, dish_id, drink_id, rating, comment) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			review.UserID, review.DishID, review.DrinkID, review.Rating, review.Comment).Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		review.ID = id
		json.NewEncoder(w).Encode(review)
	}
}

// UpdateReview

func UpdateReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var review models.Review
		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if review.ID == 0 {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE reviews SET user_id = $1, dish_id = $2, drink_id = $3, rating = $4, comment = $5 WHERE id = $6",
			review.UserID, review.DishID, review.DrinkID, review.Rating, review.Comment, review.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(review)
	}
}

// DeleteReview

func DeleteReview(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("DELETE FROM reviews WHERE id = $1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
