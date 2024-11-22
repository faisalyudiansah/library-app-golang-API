package services

import (
	"context"

	apperrors "library-api/appErrors"
	"library-api/dtos"
	"library-api/repositories"
)

type BookService interface {
	GetAllBookService(context.Context, string) ([]dtos.ResponseBook, error)
	PostBookService(context.Context, dtos.RequestPostBook) (dtos.ResponseBook, error)
}

type BookServiceImplementation struct {
	BookRepository        repositories.BookRepository
	AuthorRepository      repositories.AuthorRepository
	BorrowRepository      repositories.BorrowRepository
	UserRepository        repositories.UserRepository
	TransactionRepository repositories.TransactionRepository
}

func NewBookServiceImplementation(
	br repositories.BookRepository,
	ar repositories.AuthorRepository,
	brr repositories.BorrowRepository,
	ur repositories.UserRepository,
	dc repositories.TransactionRepository,
) *BookServiceImplementation {
	return &BookServiceImplementation{
		BookRepository:        br,
		AuthorRepository:      ar,
		BorrowRepository:      brr,
		UserRepository:        ur,
		TransactionRepository: dc,
	}
}

func (bsi *BookServiceImplementation) GetAllBookService(ctx context.Context, query string) ([]dtos.ResponseBook, error) {
	books, err := bsi.BookRepository.GetAllRepository(ctx, query)
	if err != nil {
		return nil, apperrors.ErrISE
	}
	return dtos.ToResponseBookAuthor(books), nil
}

func (bsi *BookServiceImplementation) PostBookService(ctx context.Context, reqBody dtos.RequestPostBook) (dtos.ResponseBook, error) {
	if ok := bsi.BookRepository.IsBookHasTheSameTitle(ctx, reqBody.Title); !ok {
		return dtos.ResponseBook{}, apperrors.ErrTitleBookAlreadyExists
	}
	getAuthor, err := bsi.AuthorRepository.GetAuthorById(ctx, reqBody.AuthorId)
	if err != nil {
		return dtos.ResponseBook{}, apperrors.ErrInvalidAuthorId
	}
	bookData, err := bsi.BookRepository.PostBookRepository(ctx, reqBody, *getAuthor)
	if err != nil {
		return dtos.ResponseBook{}, apperrors.ErrFailedAddProduct
	}
	return dtos.ToDtoResponseBookAuthor(*bookData), nil
}
