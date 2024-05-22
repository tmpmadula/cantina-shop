// internal/db/db.go
package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	// Create users table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT,
            email TEXT,
			password TEXT,
			role TEXT
        )
    `)
	if err != nil {
		return nil, err
	}

	// Create dishes table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS dishes (
            id SERIAL PRIMARY KEY,
            name TEXT,
            description TEXT,
            price DECIMAL(10, 2),
            image TEXT,
            rating DECIMAL(3, 2)
        )
    `)
	if err != nil {
		return nil, err
	}

	// Create drinks table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS drinks (
            id SERIAL PRIMARY KEY,
            name TEXT,
            description TEXT,
            price DECIMAL(10, 2),
            image TEXT,
            rating DECIMAL(3, 2)
        )
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}
