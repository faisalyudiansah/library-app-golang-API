package dtos

import (
	"time"
)

type ResponseMessageOnly struct {
	Message string `json:"message"`
}

type ResponseManyData struct {
	Message   string                    `json:"message"`
	TotalData int64                     `json:"totalData"`
	Data      map[string][]ResponseBook `json:"data"`
}

type ResponseOneData struct {
	Message   string                  `json:"message"`
	TotalData int64                   `json:"totalData"`
	Data      map[string]ResponseBook `json:"data"`
}

type ResponseOneDataBorrow struct {
	Message string                    `json:"message"`
	Result  map[string]ResponseBorrow `json:"result"`
}

type ResponseRegisterUser struct {
	Message string                      `json:"message"`
	Result  map[string]ResponseDataUser `json:"result"`
}

type ResponseBook struct {
	ID          int64          `json:"id"`
	AuthorId    int64          `json:"author_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Quantity    int64          `json:"quantity"`
	Cover       *string        `json:"cover"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeleteAt    *time.Time     `json:"deleted_at"`
	Author      ResponseAuthor `json:"author"`
}

type ResponseAuthor struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeleteAt  *time.Time `json:"deleted_at"`
}

type ResponseApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

type ResponseBorrow struct {
	ID         int64      `json:"id"`
	UserID     int64      `json:"user_id"`
	BookID     int64      `json:"book_id"`
	BorrowDate time.Time  `json:"borrow_date"`
	ReturnDate *time.Time `json:"return_date"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeleteAt   *time.Time `json:"deleted_at"`
}

type ResponseDataUser struct {
	AccessToken *string    `json:"access_token,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Email       *string    `json:"email,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
}
