package db

import (
	"github.com/tmpmadula/cantina-shop/internal/models"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func InitDB() {
	db = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})

	// Create schema for User model
	err := db.Model((*models.User)(nil)).CreateTable(&pg.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := db.Model(user).Where("email = ?", email).First()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *models.User) error {
	_, err := db.Model(user).Insert()
	return err
}
