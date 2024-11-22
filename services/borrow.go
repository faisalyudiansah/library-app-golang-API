package services

import (
	"context"

	apperrors "library-api/appErrors"
	"library-api/dtos"
	"library-api/models"
	"library-api/repositories"
)

type BorrowService interface {
	PostNewBorrowService(context.Context, dtos.RequestBorrowBook, int64) (*dtos.ResponseBorrow, error)
	PostReturnBookService(context.Context, dtos.RequestReturnBook, int64) (*dtos.ResponseBorrow, error)
}

type BorrowServiceImplementation struct {
	BookRepository        repositories.BookRepository
	BorrowRepository      repositories.BorrowRepository
	UserRepository        repositories.UserRepository
	TransactionRepository repositories.TransactionRepository
}

func NewBorrowServiceImplementation(
	br repositories.BookRepository,
	brr repositories.BorrowRepository,
	ur repositories.UserRepository,
	dc repositories.TransactionRepository,
) *BorrowServiceImplementation {
	return &BorrowServiceImplementation{
		BookRepository:        br,
		BorrowRepository:      brr,
		UserRepository:        ur,
		TransactionRepository: dc,
	}
}

func (bsi *BorrowServiceImplementation) PostNewBorrowService(ctx context.Context, reqBody dtos.RequestBorrowBook, userId int64) (*dtos.ResponseBorrow, error) {
	result, err := bsi.TransactionRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		findUser, err := bsi.UserRepository.GetUserById(cForTx, userId)
		if err != nil {
			return nil, apperrors.ErrInvalidAccessToken
		}
		findBook, err := bsi.BookRepository.GetBookByID(cForTx, reqBody.BookID)
		if err != nil {
			return nil, apperrors.ErrInvalidBookId
		}
		if isBorrowNow := bsi.BorrowRepository.IsUserBorrowNow(cForTx, findUser.ID); isBorrowNow {
			return nil, apperrors.ErrUserAlreadyBorrowBook
		}
		if bsi.BookRepository.IsBookOutOfStock(findBook.Quantity) {
			return nil, apperrors.ErrOutOfStockBook
		}
		result, err := bsi.BorrowRepository.PostNewBorrow(cForTx, reqBody, findUser.ID)
		if err != nil {
			return nil, apperrors.ErrFailedBorrowBook
		}
		err = bsi.BookRepository.PutQuantityBook(cForTx, int(findBook.Quantity-1), findBook.ID)
		if err != nil {
			return nil, apperrors.ErrFailedBorrowBook
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return dtos.ToDtoResponseBorrow(result.(*models.Borrow)), nil
}

func (bsi *BorrowServiceImplementation) PostReturnBookService(ctx context.Context, reqBody dtos.RequestReturnBook, userId int64) (*dtos.ResponseBorrow, error) {
	result, err := bsi.TransactionRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		findUser, err := bsi.UserRepository.GetUserById(cForTx, userId)
		if err != nil {
			return nil, apperrors.ErrInvalidAccessToken
		}
		findBook, err := bsi.BookRepository.GetBookByID(cForTx, reqBody.BookID)
		if err != nil {
			return nil, apperrors.ErrInvalidBookId
		}
		if IsReqIdValid := bsi.BorrowRepository.IsBorrowIdValid(cForTx, reqBody.BorrowID, findBook.ID); !IsReqIdValid {
			return nil, apperrors.ErrInvalidBorrowIdBookId
		}
		if isReturnBefore := bsi.BorrowRepository.IsAlreadyReturnBook(cForTx, reqBody.BorrowID, findUser.ID, findBook.ID); isReturnBefore {
			return nil, apperrors.ErrUserAlreadyReturnBook
		}
		result, err := bsi.BorrowRepository.PostReturnBook(cForTx, reqBody)
		if err != nil {
			return nil, apperrors.ErrFailedReturnBook
		}
		err = bsi.BookRepository.PutQuantityBook(cForTx, int(findBook.Quantity+1), findBook.ID)
		if err != nil {
			return nil, apperrors.ErrFailedReturnBook
		}
		return result, nil
	})
	if err != nil {
		return nil, err
	}
	return dtos.ToDtoResponseBorrow(result.(*models.Borrow)), nil
}
