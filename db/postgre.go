package db

import (
	"abcShop/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewSqlConnection(user string, password string, dbName string) *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", user, password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Product{}, &models.User{})

	return db
}
