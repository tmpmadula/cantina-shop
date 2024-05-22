package db

import (
	"errors"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/tmpmadula/cantina-shop/internal/models"
)

var db *pg.DB

func InitDB() {
	db = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "user",
		Password: "password",
		Database: "your_db_name",
	})

	// Create schema for User and Login models

	err := db.Model((*models.User)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		panic(err)
	}

	err = db.Model((*models.Login)(nil)).CreateTable(&orm.CreateTableOptions{
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
		if err == pg.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func CreateUser(user *models.User) error {
	_, err := db.Model(user).Insert()
	return err
}

func SaveLogin(userID int64, method string) error {
	login := &models.Login{
		UserID:    userID,
		Method:    method,
		Timestamp: time.Now().String(),
	}
	_, err := db.Model(login).Insert()
	return err
}
