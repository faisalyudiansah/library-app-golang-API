package dtos

import "time"

type RequestPostBook struct {
	Title       string  `json:"title" binding:"required"`
	AuthorId    int64   `json:"author_id" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Quantity    int64   `json:"quantity" binding:"required"`
	Cover       *string `json:"cover,omitempty"`
}

type RequestBorrowBook struct {
	BookID     int64      `json:"book_id" binding:"required"`
	BorrowDate *time.Time `json:"borrow_date,omitempty"`
}

type RequestReturnBook struct {
	BorrowID   int64      `json:"borrow_id" binding:"required"`
	BookID     int64      `json:"book_id" binding:"required"`
	ReturnDate *time.Time `json:"return_date" binding:"required"`
}

type RequestValidationMiddleware struct {
	Title       string     `json:"title" binding:"required"`
	AuthorId    int64      `json:"author_id" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Quantity    int64      `json:"quantity" binding:"required"`
	BookID      int64      `json:"book_id" binding:"required"`
	BorrowID    int64      `json:"borrow_id" binding:"required"`
	ReturnDate  *time.Time `json:"return_date" binding:"required"`
	Name        string     `json:"name" binding:"required"`
	Email       string     `json:"email" binding:"required"`
	Password    string     `json:"password" binding:"required"`
}

type RequestRegisterUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RequestLoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
