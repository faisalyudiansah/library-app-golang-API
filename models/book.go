package models

import "time"

type Book struct {
	ID          int64
	AuthorId    int64
	Title       string
	Description string
	Quantity    int64
	Cover       *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeleteAt    *time.Time
}

type AuthorBook struct {
	Book   Book
	Author Author
}
