// internal/models/review.go
package models

type Review struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	DishID  int    `json:"dish_id"`
	DrinkID int    `json:"drink_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment,omitempty"`
}
