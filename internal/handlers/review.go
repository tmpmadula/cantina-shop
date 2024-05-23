// internal/handlers/review.go
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
