package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BookReading struct {
	gorm.Model
	BookID    int
	Book      Book
	UserID    int
	User      User
	StartDate datatypes.Date
}
