// cmd/main.go
package main

import (
	"github.com/tmpmadula/cantina-shop/internal/api"
	"github.com/tmpmadula/cantina-shop/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database
	db.InitDB()

	// Initialize API routes
	api.InitRoutes(r)

	r.Run(":8080")
}
