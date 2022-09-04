package entities

import "gorm.io/gorm"

type Bookshelf struct {
	gorm.Model
	UserID   int
	User     User
	Name     string
	Location string
	Tag      string
	Books    []Book
}
