package services_test

import (
	"context"
	"testing"
	"time"

	apperrors "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/appErrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/services"
	"github.com/stretchr/testify/assert"
)

func TestBookServiceImplementation_GetAllBookService(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		mockRepository func(*mocks.BookRepository, []any)
		want           []models.AuthorBook
		err            error
		wantErr        bool
	}{
		{
			name:  "should success get all books when the book service is active",
			query: "Frieren",
			mockRepository: func(mock *mocks.BookRepository, result []any) {
				mock.On("GetAllRepository", context.Background(), "Frieren").Return(result...)
			},
			want: []models.AuthorBook{
				{
					Book:   models.Book{ID: 1, Title: "Frieren rambut lurus"},
					Author: models.Author{ID: 1, Name: "Frieren"},
				},
				{
					Book:   models.Book{ID: 1, Title: "Batman rambut lurus"},
					Author: models.Author{ID: 1, Name: "Batman"},
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name:  "should return error when get all service is active",
			query: "Frieren",
			mockRepository: func(mock *mocks.BookRepository, result []any) {
				mock.On("GetAllRepository", context.Background(), "Frieren").Return(result...)
			},
			want:    nil,
			err:     apperrors.ErrISE,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := &mocks.BookRepository{}
			mockAuthor := &mocks.AuthorRepository{}
			mockBorrow := &mocks.BorrowRepository{}
			mockUser := &mocks.UserRepository{}
			mockTx := &mocks.TransactionRepository{}
			bookService := services.NewBookServiceImplementation(mockBook, mockAuthor, mockBorrow, mockUser, mockTx)
			tt.mockRepository(mockBook, []any{tt.want, tt.err})

			books, err := bookService.GetAllBookService(context.Background(), tt.query)

			if tt.wantErr {
				assert.Nil(t, books)
				assert.Equal(t, apperrors.ErrISE, err)
			} else {
				assert.Equal(t, tt.err, err)
				assert.Equal(t, dtos.ToResponseBookAuthor(tt.want), books)
			}
		})
	}
}

func TestBookServiceImplementation_PostBookService(t *testing.T) {
	book := models.Book{
		ID:          1,
		AuthorId:    1,
		Title:       "Rasenggan",
		Description: "Colokan",
		Quantity:    5,
		Cover:       nil,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeleteAt:    nil,
	}
	author := &models.Author{
		ID:        1,
		Name:      "Ronaldo",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeleteAt:  nil,
	}
	bookAuthor := &models.AuthorBook{
		Book:   book,
		Author: *author,
	}
	tests := []struct {
		name           string
		reqBody        dtos.RequestPostBook
		mockRepository func(*mocks.BookRepository, *mocks.AuthorRepository, dtos.RequestPostBook)
		want           dtos.ResponseBook
		err            error
		wantErr        bool
	}{
		{
			name: "should successfully add a new book with service book",
			reqBody: dtos.RequestPostBook{
				Title:       "Rasenggan",
				AuthorId:    1,
				Description: "Colokan",
				Quantity:    5,
				Cover:       nil,
			},
			mockRepository: func(mockBook *mocks.BookRepository, mockAuthor *mocks.AuthorRepository, reqBody dtos.RequestPostBook) {
				mockBook.On("IsBookHasTheSameTitle", context.Background(), reqBody.Title).Return(true)
				mockAuthor.On("GetAuthorById", context.Background(), reqBody.AuthorId).Return(author, nil)
				mockBook.On("PostBookRepository", context.Background(), reqBody, *author).Return(bookAuthor, nil)
			},
			want: dtos.ResponseBook{
				ID:          1,
				Title:       "Rasenggan",
				AuthorId:    1,
				Description: "Colokan",
				Quantity:    5,
				Cover:       nil,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
				DeleteAt:    nil,
				Author: dtos.ResponseAuthor{
					ID:        1,
					Name:      "Ronaldo",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
					DeleteAt:  nil,
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "should error when add a new book with the same title of book",
			reqBody: dtos.RequestPostBook{
				Title:       "Telor sama sayur ga pedes",
				AuthorId:    1,
				Description: "Cerita rakyat anak pendalaman",
				Quantity:    5,
				Cover:       nil,
			},
			mockRepository: func(mockBook *mocks.BookRepository, mockAuthor *mocks.AuthorRepository, reqBody dtos.RequestPostBook) {
				mockBook.On("IsBookHasTheSameTitle", context.Background(), reqBody.Title).Return(false)
			},
			want:    dtos.ResponseBook{},
			err:     apperrors.ErrTitleBookAlreadyExists,
			wantErr: true,
		},
		{
			name:    "should error when user sent invalid req body",
			reqBody: dtos.RequestPostBook{},
			mockRepository: func(mockBook *mocks.BookRepository, mockAuthor *mocks.AuthorRepository, reqBody dtos.RequestPostBook) {
				mockBook.On("IsBookHasTheSameTitle", context.Background(), reqBody.Title).Return(true)
				mockAuthor.On("GetAuthorById", context.Background(), reqBody.AuthorId).Return(author, nil)
				mockBook.On("PostBookRepository", context.Background(), reqBody, *author).Return(nil, apperrors.ErrFailedAddProduct)
			},
			want:    dtos.ResponseBook{},
			err:     apperrors.ErrFailedAddProduct,
			wantErr: true,
		},
		{
			name:    "should error when user sent invalid author id",
			reqBody: dtos.RequestPostBook{},
			mockRepository: func(mockBook *mocks.BookRepository, mockAuthor *mocks.AuthorRepository, reqBody dtos.RequestPostBook) {
				mockBook.On("IsBookHasTheSameTitle", context.Background(), reqBody.Title).Return(true)
				mockAuthor.On("GetAuthorById", context.Background(), reqBody.AuthorId).Return(nil, apperrors.ErrInvalidAuthorId)
			},
			want:    dtos.ResponseBook{},
			err:     apperrors.ErrInvalidAuthorId,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBook := &mocks.BookRepository{}
			mockAuthor := &mocks.AuthorRepository{}
			mockBorrow := &mocks.BorrowRepository{}
			mockUser := &mocks.UserRepository{}
			mockTx := &mocks.TransactionRepository{}
			bookService := services.NewBookServiceImplementation(mockBook, mockAuthor, mockBorrow, mockUser, mockTx)
			tt.mockRepository(mockBook, mockAuthor, tt.reqBody)

			book, err := bookService.PostBookService(context.Background(), tt.reqBody)

			if tt.wantErr {
				assert.Equal(t, tt.err, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, book)
			}
		})
	}
}
