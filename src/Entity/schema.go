package Entity

import "time"

type Book struct {
	Id        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"'`
	Author    string    `json:"author" db:"author"`
	Price     int       `json:"price" db:"price"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
