package controllers

import (
	"io"
	"net/http"

	apperrors "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/appErrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/constants"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/services"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService services.BookService
}

func NewBookController(bs services.BookService) *BookController {
	return &BookController{
		BookService: bs,
	}
}

func (bc *BookController) GetAllBookController(c *gin.Context) {
	query := c.Query("title")
	books, err := bc.BookService.GetAllBookService(c, query)
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterManyBook(books, constants.Ok))
}

func (bc *BookController) PostBookController(c *gin.Context) {
	reqBody := dtos.RequestPostBook{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		if err == io.EOF {
			c.Error(apperrors.ErrRequestBodyInvalid)
			return
		}
		c.Error(err)
		return
	}
	book, err := bc.BookService.PostBookService(c, reqBody)
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusCreated, helpers.FormatterOneBook(book, constants.SuccessAddNewBook))
}
