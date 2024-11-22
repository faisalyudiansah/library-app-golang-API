package models

import "time"

type Borrow struct {
	ID         int64
	UserID     int64
	BookID     int64
	BorrowDate time.Time
	ReturnDate *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   *time.Time
}

type NewBorrow struct {
	UserID     int64
	BookID     int64
	BorrowDate time.Time
}
