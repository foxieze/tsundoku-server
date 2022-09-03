package config

import (
	"github.com/foxieze/tsundoku-api/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:root@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.User{}, &entities.Bookshelf{}, &entities.Book{}, &entities.BookReading{}, &entities.BookRead{})

	return nil
}