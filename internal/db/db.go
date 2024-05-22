package db

import (
	"errors"
	"time"

	"log"

	"github.com/tmpmadula/cantina-shop/config"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/tmpmadula/cantina-shop/internal/models"
)

var db *pg.DB

func InitDB() error {
	connStr := config.GetDBConnectionString()
	db = pg.Connect(&pg.Options{
		Addr: connStr,

		User:     "postgres", // Assuming you have a default user
		Password: "postgres", // Assuming you have a default password
		Database: "postgres",
	})

	// Perform a simple query to check the connection
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		return errors.New("failed to connect to database: " + err.Error())
	}

	log.Println("Database connection established")
	return nil
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
	// check if table exists, if not create it
	// console.log("Creating table")
	log.Print("Creating table")
	err := db.Model(user).CreateTable(&orm.CreateTableOptions{
		Temp: false,
	})
	if err != nil {
		return err
	}

	// insert user
	_, err = db.Model(user).Insert()
	return err

	//_, err := db.Model(user).Insert()
	//return err
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

func GetAllDishes() ([]*models.Dish, error) {
	var dishes []*models.Dish
	err := db.Model(&dishes).Select()
	return dishes, err
}
