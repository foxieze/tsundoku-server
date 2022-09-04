package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookshelfID int
	Bookshelf   Bookshelf
	ISBN        int
	Title       string
	Author      string
}
