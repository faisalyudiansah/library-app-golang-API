package apperrors

import (
	"errors"

	"library-api/constants"
)

var (
	ErrTitleBookAlreadyExists = errors.New(constants.FailedTitleAlreadyExists)
	ErrFailedAddProduct       = errors.New(constants.FailedAddProduct)
	ErrRequestBodyInvalid     = errors.New(constants.RequestBodyInvalid)
	ErrInvalidAuthorId        = errors.New(constants.InvalidAuthorId)
	ErrInvalidUserId          = errors.New(constants.InvalidUserId)
	ErrInvalidBookId          = errors.New(constants.InvalidBookId)
	ErrInvalidBorrowIdBookId  = errors.New(constants.InvalidBorrowIdBookId)
	ErrOutOfStockBook         = errors.New(constants.OutOfStockBook)
	ErrUserAlreadyBorrowBook  = errors.New(constants.UserAlreadyBorrowBook)
	ErrFailedBorrowBook       = errors.New(constants.FailedBorrowBook)
	ErrFailedReturnBook       = errors.New(constants.FailedReturnBook)
	ErrUserAlreadyReturnBook  = errors.New(constants.UserAlreadyReturnBook)
)

var (
	ErrUserInvalidEmailPassword = errors.New(constants.UserInvalidEmailPassword)
	ErrUserFailedRegister       = errors.New(constants.UserFailedRegister)
	ErrUserFailedLogin          = errors.New(constants.UserFailedLogin)
	ErrUserEmailAlreadyExists   = errors.New(constants.UserEmailAlreadyExists)
)

var (
	ErrISE                = errors.New(constants.ISE)
	ErrInvalidAccessToken = errors.New(constants.InvalidAccessToken)
	ErrUnauthorization    = errors.New(constants.Unauthorization)
	ErrUrlNotFound        = errors.New(constants.UrlNotFound)
)
