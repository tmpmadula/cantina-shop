// internal/api/dishes.go

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmpmadula/cantina-shop/internal/db"
)

// GetDishesHandler handles fetching all dishes
func GetDishesHandler(c *gin.Context) {
	dishes, err := db.GetAllDishes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dishes)
}
