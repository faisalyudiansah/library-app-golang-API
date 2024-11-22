package controllers

import (
	"io"
	"net/http"

	apperrors "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/appErrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/constants"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers"
	utilscontext "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers/utilsContext"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/services"
	"github.com/gin-gonic/gin"
)

type BorrowController struct {
	BorrowService services.BorrowService
}

func NewBorrowController(bs services.BorrowService) *BorrowController {
	return &BorrowController{
		BorrowService: bs,
	}
}

func (brr *BorrowController) PostNewBorrowBookController(c *gin.Context) {
	reqBody := dtos.RequestBorrowBook{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		if err == io.EOF {
			c.Error(apperrors.ErrRequestBodyInvalid)
			return
		}
		c.Error(err)
		return
	}
	data, err := brr.BorrowService.PostNewBorrowService(c, reqBody, utilscontext.GetValueUserIdFromToken(c))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusCreated, helpers.FormatterOneBorrow(data, constants.SuccessBorrowBook))
}

func (brr *BorrowController) PostReturnBookController(c *gin.Context) {
	reqBody := dtos.RequestReturnBook{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		if err == io.EOF {
			c.Error(apperrors.ErrRequestBodyInvalid)
			return
		}
		c.Error(err)
		return
	}
	result, err := brr.BorrowService.PostReturnBookService(c, reqBody, utilscontext.GetValueUserIdFromToken(c))
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterOneBorrow(result, constants.SuccessReturnBook))
}
