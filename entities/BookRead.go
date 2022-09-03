package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type BookRead struct {
	gorm.Model
	BookID    int
	Book      Book
	UserID    int
	User      User
	StartDate datatypes.Date
	EndDate   datatypes.Date
}
