// internal/db/database.go
package db

import (
	"github.com/go-pg/pg"
)

var db *pg.DB

func InitDB() {
	// Initialize database connection
	// Use config package to load database configurations
}

// Add functions to interact with the database
