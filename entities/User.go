package entities

type User struct {
	Id               int `gorm:"primaryKey"`
	Username         string
	Password         string
	Name             string
	Bookshelves      []Bookshelf
	CurrentlyReading []BookReading
	FinishedReading  []BookRead
}
