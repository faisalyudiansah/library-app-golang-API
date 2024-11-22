package controllers

import (
	"io"
	"net/http"

	apperrors "library-api/appErrors"
	"library-api/constants"
	"library-api/dtos"
	"library-api/helpers"
	"library-api/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(bs services.UserService) *UserController {
	return &UserController{
		UserService: bs,
	}
}

func (uc *UserController) PostRegisterUserController(c *gin.Context) {
	reqBody := dtos.RequestRegisterUser{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		if err == io.EOF {
			c.Error(apperrors.ErrRequestBodyInvalid)
			return
		}
		c.Error(err)
		return
	}
	dataUser, err := uc.UserService.PostRegisterUserService(c, reqBody)
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusCreated, helpers.FormatterSuccessRegisterLogin(dataUser, constants.SuccessRegister))
}

func (uc *UserController) PostLoginUserController(c *gin.Context) {
	reqBody := dtos.RequestLoginUser{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		if err == io.EOF {
			c.Error(apperrors.ErrRequestBodyInvalid)
			return
		}
		c.Error(err)
		return
	}
	dataUser, err := uc.UserService.PostLoginUserService(c, reqBody)
	if err != nil {
		c.Error(err)
		return
	}
	helpers.PrintResponse(c, http.StatusOK, helpers.FormatterSuccessRegisterLogin(dataUser, constants.SuccessLogin))
}
