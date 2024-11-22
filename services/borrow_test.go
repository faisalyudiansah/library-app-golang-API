package services

import (
	"context"
	"testing"
	"time"

	apperrors "library-api/appErrors"
	"library-api/dtos"
	"library-api/mocks"
	"library-api/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userId       = int64(1)
	bookId       = int64(2)
	nameUser     = "Frierennnnnnn"
	emailUser    = "frieren@gmail.com"
	passwordUser = "12345"
)
var resUser *models.User = &models.User{
	ID:        1,
	Name:      nameUser,
	Email:     emailUser,
	Password:  passwordUser,
	CreatedAt: time.Time{},
	UpdatedAt: &time.Time{},
	DeleteAt:  nil,
}
var resBook *models.Book = &models.Book{
	ID:          1,
	AuthorId:    1,
	Title:       "FRIERENNNNNNNNNNNNNNNNNNNNN",
	Description: "NONA FRIEREN BAGUS",
	Quantity:    5,
	Cover:       nil,
	CreatedAt:   time.Time{},
	UpdatedAt:   time.Time{},
	DeleteAt:    nil,
}
var resBorrow *models.Borrow = &models.Borrow{
	ID:         int64(1),
	UserID:     int64(1),
	BookID:     int64(2),
	BorrowDate: time.Time{},
	ReturnDate: nil,
	CreatedAt:  time.Time{},
	UpdatedAt:  time.Time{},
	DeleteAt:   nil,
}
var resReturn *models.Borrow = &models.Borrow{
	ID:         int64(1),
	UserID:     int64(1),
	BookID:     int64(2),
	BorrowDate: time.Time{},
	ReturnDate: &time.Time{},
	CreatedAt:  time.Time{},
	UpdatedAt:  time.Time{},
	DeleteAt:   nil,
}

func TestBorrowServiceImplementation_PostNewBorrowService(t *testing.T) {
	tests := []struct {
		name           string
		mockRepository func(*mocks.TransactionRepository, *mocks.BookRepository, *mocks.UserRepository, *mocks.BorrowRepository, dtos.RequestBorrowBook)
		reqBody        dtos.RequestBorrowBook
		want           *dtos.ResponseBorrow
		wantErr        error
	}{
		{
			name: "success borrow a book when user input the correct request",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				reqBody dtos.RequestBorrowBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(resBorrow, nil)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, reqBody.BookID).Return(resBook, nil)
				br2.On("IsUserBorrowNow", mock.Anything, resUser.ID).Return(false)
				br1.On("IsBookOutOfStock", resBook.Quantity).Return(false)
				br2.On("PostNewBorrow", mock.Anything, reqBody, resUser.ID).Return(resBorrow, nil)
				br1.On("PutQuantityBook", mock.Anything, int(resBook.Quantity-1), resBook.ID).Return(nil)
			},
			reqBody: dtos.RequestBorrowBook{
				BookID:     bookId,
				BorrowDate: &time.Time{},
			},
			want: &dtos.ResponseBorrow{
				ID:         resBorrow.ID,
				UserID:     resBorrow.UserID,
				BookID:     resBorrow.BookID,
				BorrowDate: resBorrow.BorrowDate,
				ReturnDate: resBorrow.ReturnDate,
				CreatedAt:  resBorrow.CreatedAt,
				UpdatedAt:  resBorrow.UpdatedAt,
				DeleteAt:   resBorrow.DeleteAt,
			},
			wantErr: nil,
		},
		{
			name: "should error when user does not registered on system",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				reqBody dtos.RequestBorrowBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrInvalidAccessToken)
				ur.On("GetUserById", mock.Anything, resUser.ID).Return(nil, apperrors.ErrInvalidAccessToken)
			},
			reqBody: dtos.RequestBorrowBook{
				BookID:     bookId,
				BorrowDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrInvalidAccessToken,
		},
		{
			name: "should error when book does not listed on system",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				reqBody dtos.RequestBorrowBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrInvalidBookId)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, reqBody.BookID).Return(nil, apperrors.ErrInvalidBookId)
			},
			reqBody: dtos.RequestBorrowBook{
				BookID:     bookId,
				BorrowDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrInvalidBookId,
		},
		{
			name: "should error when user already borrow the book",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				reqBody dtos.RequestBorrowBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrUserAlreadyBorrowBook)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, reqBody.BookID).Return(resBook, nil)
				br2.On("IsUserBorrowNow", mock.Anything, resUser.ID).Return(true)
			},
			reqBody: dtos.RequestBorrowBook{
				BookID:     bookId,
				BorrowDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrUserAlreadyBorrowBook,
		},
		{
			name: "should error when quantity / stock book is out of stock",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				reqBody dtos.RequestBorrowBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrOutOfStockBook)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, reqBody.BookID).Return(resBook, nil)
				br2.On("IsUserBorrowNow", mock.Anything, resUser.ID).Return(false)
				br1.On("IsBookOutOfStock", resBook.Quantity).Return(true)
			},
			reqBody: dtos.RequestBorrowBook{
				BookID:     bookId,
				BorrowDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrOutOfStockBook,
		},
		{
			name: "should error when system want to save new borrow but something wrong in system or in db",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				reqBody dtos.RequestBorrowBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrFailedBorrowBook)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, reqBody.BookID).Return(resBook, nil)
				br2.On("IsUserBorrowNow", mock.Anything, resUser.ID).Return(false)
				br1.On("IsBookOutOfStock", resBook.Quantity).Return(false)
				br2.On("PostNewBorrow", mock.Anything, reqBody, resUser.ID).Return(nil, apperrors.ErrFailedBorrowBook)
			},
			reqBody: dtos.RequestBorrowBook{
				BookID:     bookId,
				BorrowDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrFailedBorrowBook,
		},
		{
			name: "should error when something wrong in repository to change the quantity of book",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				reqBody dtos.RequestBorrowBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrFailedBorrowBook)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, reqBody.BookID).Return(resBook, nil)
				br2.On("IsUserBorrowNow", mock.Anything, resUser.ID).Return(false)
				br1.On("IsBookOutOfStock", resBook.Quantity).Return(false)
				br2.On("PostNewBorrow", mock.Anything, reqBody, resUser.ID).Return(resBorrow, nil)
				br1.On("PutQuantityBook", mock.Anything, int(resBook.Quantity-1), resBook.ID).Return(apperrors.ErrFailedBorrowBook)
			},
			reqBody: dtos.RequestBorrowBook{
				BookID:     bookId,
				BorrowDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrFailedBorrowBook,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := &mocks.BookRepository{}
			mockBorrow := &mocks.BorrowRepository{}
			mockUser := &mocks.UserRepository{}
			mockTx := &mocks.TransactionRepository{}
			borrowService := NewBorrowServiceImplementation(mockBook, mockBorrow, mockUser, mockTx)
			tt.mockRepository(mockTx, mockBook, mockUser, mockBorrow, tt.reqBody)

			result, err := borrowService.PostNewBorrowService(context.Background(), tt.reqBody, userId)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestBorrowServiceImplementation_PostReturnBookService(t *testing.T) {
	tests := []struct {
		name           string
		mockRepository func(*mocks.TransactionRepository, *mocks.BookRepository, *mocks.UserRepository, *mocks.BorrowRepository, dtos.RequestReturnBook)
		reqBody        dtos.RequestReturnBook
		want           *dtos.ResponseBorrow
		wantErr        error
	}{
		{
			name: "should successfully when user want to return a book to library",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				rrb dtos.RequestReturnBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(resReturn, nil)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, rrb.BookID).Return(resBook, nil)
				br2.On("IsBorrowIdValid", mock.Anything, rrb.BookID, resBook.ID).Return(true)
				br2.On("IsAlreadyReturnBook", mock.Anything, rrb.BorrowID, userId, rrb.BookID).Return(false)
				br2.On("PostReturnBook", mock.Anything, rrb).Return(resReturn, nil)
				br1.On("PutQuantityBook", mock.Anything, int(resBook.Quantity+1), resBook.ID).Return(nil)
			},
			reqBody: dtos.RequestReturnBook{
				BorrowID:   int64(1),
				BookID:     int64(1),
				ReturnDate: &time.Time{},
			},
			want: &dtos.ResponseBorrow{
				ID:         resReturn.ID,
				UserID:     resReturn.UserID,
				BookID:     resReturn.BookID,
				BorrowDate: resReturn.BorrowDate,
				ReturnDate: resReturn.ReturnDate,
				CreatedAt:  resReturn.CreatedAt,
				UpdatedAt:  resReturn.UpdatedAt,
				DeleteAt:   resReturn.DeleteAt,
			},
			wantErr: nil,
		},
		{
			name: "should error when user does not registered on system",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				rrb dtos.RequestReturnBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrInvalidAccessToken)
				ur.On("GetUserById", mock.Anything, userId).Return(nil, apperrors.ErrInvalidAccessToken)
			},
			reqBody: dtos.RequestReturnBook{
				BorrowID:   int64(1),
				BookID:     int64(1),
				ReturnDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrInvalidAccessToken,
		},
		{
			name: "should error when book does not listed on system",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				rrb dtos.RequestReturnBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrInvalidBookId)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, rrb.BookID).Return(nil, apperrors.ErrInvalidBookId)
			},
			reqBody: dtos.RequestReturnBook{
				BorrowID:   int64(1),
				BookID:     int64(1),
				ReturnDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrInvalidBookId,
		},
		{
			name: "should error when borrow id does not listed on system",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				rrb dtos.RequestReturnBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrInvalidBorrowIdBookId)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, rrb.BookID).Return(resBook, nil)
				br2.On("IsBorrowIdValid", mock.Anything, rrb.BookID, resBook.ID).Return(false)
			},
			reqBody: dtos.RequestReturnBook{
				BorrowID:   int64(1),
				BookID:     int64(1),
				ReturnDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrInvalidBorrowIdBookId,
		},
		{
			name: "should error when user already return a book",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				rrb dtos.RequestReturnBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrInvalidBorrowIdBookId)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, rrb.BookID).Return(resBook, nil)
				br2.On("IsBorrowIdValid", mock.Anything, rrb.BookID, resBook.ID).Return(true)
				br2.On("IsAlreadyReturnBook", mock.Anything, rrb.BorrowID, userId, rrb.BookID).Return(true)
			},
			reqBody: dtos.RequestReturnBook{
				BorrowID:   int64(1),
				BookID:     int64(1),
				ReturnDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrInvalidBorrowIdBookId,
		},
		{
			name: "should be an error when the book return process in the database system fails",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				rrb dtos.RequestReturnBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrFailedReturnBook)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, rrb.BookID).Return(resBook, nil)
				br2.On("IsBorrowIdValid", mock.Anything, rrb.BookID, resBook.ID).Return(true)
				br2.On("IsAlreadyReturnBook", mock.Anything, rrb.BorrowID, userId, rrb.BookID).Return(false)
				br2.On("PostReturnBook", mock.Anything, rrb).Return(nil, apperrors.ErrFailedReturnBook)
			},
			reqBody: dtos.RequestReturnBook{
				BorrowID:   int64(1),
				BookID:     int64(1),
				ReturnDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrFailedReturnBook,
		},
		{
			name: "should error when something wrong in repository to change the quantity of book",
			mockRepository: func(
				tr *mocks.TransactionRepository,
				br1 *mocks.BookRepository,
				ur *mocks.UserRepository,
				br2 *mocks.BorrowRepository,
				rrb dtos.RequestReturnBook,
			) {
				tr.On("Atomic", mock.Anything, mock.MatchedBy(func(callback func(context.Context) (any, error)) bool {
					callback(context.Background())
					return true
				})).Return(nil, apperrors.ErrFailedReturnBook)
				ur.On("GetUserById", mock.Anything, userId).Return(resUser, nil)
				br1.On("GetBookByID", mock.Anything, rrb.BookID).Return(resBook, nil)
				br2.On("IsBorrowIdValid", mock.Anything, rrb.BookID, resBook.ID).Return(true)
				br2.On("IsAlreadyReturnBook", mock.Anything, rrb.BorrowID, userId, rrb.BookID).Return(false)
				br2.On("PostReturnBook", mock.Anything, rrb).Return(resReturn, nil)
				br1.On("PutQuantityBook", mock.Anything, int(resBook.Quantity+1), resBook.ID).Return(apperrors.ErrFailedReturnBook)
			},
			reqBody: dtos.RequestReturnBook{
				BorrowID:   int64(1),
				BookID:     int64(1),
				ReturnDate: &time.Time{},
			},
			want:    nil,
			wantErr: apperrors.ErrFailedReturnBook,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTx := &mocks.TransactionRepository{}
			mockBook := &mocks.BookRepository{}
			mockBorrow := &mocks.BorrowRepository{}
			mockUser := &mocks.UserRepository{}
			borrowService := NewBorrowServiceImplementation(mockBook, mockBorrow, mockUser, mockTx)
			tt.mockRepository(mockTx, mockBook, mockUser, mockBorrow, tt.reqBody)

			result, err := borrowService.PostReturnBookService(context.Background(), tt.reqBody, userId)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
