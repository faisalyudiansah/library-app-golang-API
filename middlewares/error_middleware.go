package middlewares

import (
	"errors"
	"net/http"

	apperrors "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/appErrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) == 0 {
		return
	}
	if len(c.Errors) > 0 {
		var ve validator.ValidationErrors
		if errors.As(c.Errors[0].Err, &ve) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": helpers.FormatterErrorInput(ve)})
			return
		}

		errorMappings := map[error]int{
			apperrors.ErrISE:                      http.StatusInternalServerError,
			apperrors.ErrRequestBodyInvalid:       http.StatusBadRequest,
			apperrors.ErrTitleBookAlreadyExists:   http.StatusBadRequest,
			apperrors.ErrInvalidAuthorId:          http.StatusBadRequest,
			apperrors.ErrFailedAddProduct:         http.StatusBadRequest,
			apperrors.ErrInvalidUserId:            http.StatusBadRequest,
			apperrors.ErrInvalidBookId:            http.StatusBadRequest,
			apperrors.ErrOutOfStockBook:           http.StatusBadRequest,
			apperrors.ErrUserAlreadyBorrowBook:    http.StatusBadRequest,
			apperrors.ErrFailedBorrowBook:         http.StatusBadRequest,
			apperrors.ErrFailedReturnBook:         http.StatusBadRequest,
			apperrors.ErrInvalidBorrowIdBookId:    http.StatusBadRequest,
			apperrors.ErrUserAlreadyReturnBook:    http.StatusBadRequest,
			apperrors.ErrUserFailedRegister:       http.StatusBadRequest,
			apperrors.ErrUserEmailAlreadyExists:   http.StatusBadRequest,
			apperrors.ErrUserInvalidEmailPassword: http.StatusBadRequest,
			apperrors.ErrUrlNotFound:              http.StatusNotFound,
			apperrors.ErrUnauthorization:          http.StatusUnauthorized,
			apperrors.ErrInvalidAccessToken:       http.StatusUnauthorized,
		}
		for err, statusCode := range errorMappings {
			if errors.Is(c.Errors[0].Err, err) {
				helpers.PrintError(c, statusCode, err.Error())
				return
			}
		}

		helpers.PrintError(c, http.StatusInternalServerError, apperrors.ErrISE.Error())
	}
}
