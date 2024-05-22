// internal/api/reviews.go
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewReviews(c *gin.Context) {
	// Fetch reviews from the database or any other data source
	reviews := []Review{
		{ID: 1, DishID: 1, DrinkID: 0, Rating: 4, Comment: "Great dish!"},
		{ID: 2, DishID: 0, DrinkID: 2, Rating: 5, Comment: "Awesome drink!"},
		// Add more reviews as needed
	}

	c.HTML(http.StatusOK, "reviews.html", gin.H{
		"reviews": reviews,
	})
}

type Review struct {
	ID      int    `json:"id"`
	DishID  int    `json:"dish_id"`
	DrinkID int    `json:"drink_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
