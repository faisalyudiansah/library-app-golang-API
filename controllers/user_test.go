package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	apperrors "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/appErrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/constants"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/controllers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/middlewares"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/servers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SetUpRouterUser(userController *controllers.UserController) *gin.Engine {
	h := &servers.HandlerOps{
		UserController: userController,
	}
	return servers.SetupRoute(h)
}

func TestRegisterUserControllers(t *testing.T) {
	t.Run("should success register user when user sent correct request body", func(t *testing.T) {
		mockUserService := &mocks.UserService{}

		var (
			nameBody     = "frieren"
			emailBody    = "frieren@gmail.com"
			passwordBody = "12345"
		)

		newUser := dtos.RequestRegisterUser{
			Name:     nameBody,
			Email:    emailBody,
			Password: passwordBody,
		}

		expectedUser := &dtos.ResponseDataUser{
			Name:      &nameBody,
			Email:     &emailBody,
			CreatedAt: &time.Time{},
		}
		mockUserService.On("PostRegisterUserService", mock.AnythingOfType("*gin.Context"), newUser).Return(expectedUser, nil)
		userController := controllers.NewUserController(mockUserService)

		g := SetUpRouterUser(userController)
		reqBody, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		expect := helpers.FormatterSuccessRegisterLogin(expectedUser, constants.SuccessRegister)
		expectedToJson, _ := json.Marshal(expect)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(expectedToJson), rec.Body.String())
		assert.Equal(t, expect.Message, response["message"])
	})

	t.Run("should failed register user when user does not sent correct request body", func(t *testing.T) {
		mockUserService := &mocks.UserService{}

		userController := controllers.NewUserController(mockUserService)

		g := SetUpRouterUser(userController)

		req, _ := http.NewRequest(http.MethodPost, "/register", http.NoBody)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		expected := dtos.ResponseMessageOnly{
			Message: constants.RequestBodyInvalid,
		}
		expectedToJson, _ := json.Marshal(expected)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(expectedToJson), rec.Body.String())
		assert.Equal(t, expected.Message, response["message"])
	})

	t.Run("should fail to register new user when user does not send email or any mandatory fields", func(t *testing.T) {
		mockUserService := &mocks.UserService{}
		userController := controllers.NewUserController(mockUserService)
		g := SetUpRouterUser(userController)
		g.Use(middlewares.ErrorHandler)

		newUserWithoutEmail := dtos.RequestRegisterUser{
			Name:     "Ronaldo",
			Password: "12345",
		}

		body, _ := json.Marshal(newUserWithoutEmail)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)

		expected := map[string]interface{}{
			"errors": []dtos.ResponseApiError{
				{
					Field: "email",
					Msg:   "This field is required",
				},
			},
		}

		expectedToJson, _ := json.Marshal(expected)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Nil(t, err)
		assert.NotNil(t, response["errors"])

		errors := response["errors"].([]interface{})
		errorDetail := errors[0].(map[string]interface{})

		assert.Equal(t, string(expectedToJson), rec.Body.String())
		assert.Equal(t, "email", errorDetail["field"])
		assert.Equal(t, "This field is required", errorDetail["message"])
	})

	t.Run("should failed register a new book when user send the same title book", func(t *testing.T) {
		mockUserService := &mocks.UserService{}

		var (
			nameBody     = "frieren"
			emailBody    = "frieren@gmail.com"
			passwordBody = "12345"
		)

		newUser := dtos.RequestRegisterUser{
			Name:     nameBody,
			Email:    emailBody,
			Password: passwordBody,
		}

		mockUserService.On("PostRegisterUserService", mock.AnythingOfType("*gin.Context"), newUser).Return(nil, apperrors.ErrUserEmailAlreadyExists)
		userController := controllers.NewUserController(mockUserService)

		g := SetUpRouterUser(userController)
		g.Use(middlewares.ErrorHandler)

		body, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		expected := dtos.ResponseMessageOnly{
			Message: constants.UserEmailAlreadyExists,
		}
		expectedToJson, _ := json.Marshal(expected)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(expectedToJson), rec.Body.String())
	})
}

func TestLoginUserControllers(t *testing.T) {
	t.Run("should success login user when user sent correct request body", func(t *testing.T) {
		mockUserService := &mocks.UserService{}

		var (
			emailBody    = "frieren@gmail.com"
			passwordBody = "12345"
			AccessToken  = "access_token"
		)

		newUser := dtos.RequestLoginUser{
			Email:    emailBody,
			Password: passwordBody,
		}

		expectedAccessToken := &dtos.ResponseDataUser{
			AccessToken: &AccessToken,
		}
		mockUserService.On("PostLoginUserService", mock.AnythingOfType("*gin.Context"), newUser).Return(expectedAccessToken, nil)
		userController := controllers.NewUserController(mockUserService)

		g := SetUpRouterUser(userController)

		reqBody, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		expect := helpers.FormatterSuccessRegisterLogin(expectedAccessToken, constants.SuccessLogin)
		expectedToJson, _ := json.Marshal(expect)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedToJson), rec.Body.String())
		assert.Equal(t, expect.Message, response["message"])
	})

	t.Run("should fail to login new user when user does not send email or any mandatory fields", func(t *testing.T) {
		mockUserService := &mocks.UserService{}
		userController := controllers.NewUserController(mockUserService)
		g := SetUpRouterUser(userController)
		g.Use(middlewares.ErrorHandler)

		newUserWithoutEmail := dtos.RequestRegisterUser{
			Password: "12345",
		}

		body, _ := json.Marshal(newUserWithoutEmail)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)

		expected := map[string]interface{}{
			"errors": []dtos.ResponseApiError{
				{
					Field: "email",
					Msg:   "This field is required",
				},
			},
		}

		expectedToJson, _ := json.Marshal(expected)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Nil(t, err)
		assert.NotNil(t, response["errors"])

		errors := response["errors"].([]interface{})
		errorDetail := errors[0].(map[string]interface{})

		assert.Equal(t, string(expectedToJson), rec.Body.String())
		assert.Equal(t, "email", errorDetail["field"])
		assert.Equal(t, "This field is required", errorDetail["message"])
	})

	t.Run("should failed login user when user does not sent correct request body", func(t *testing.T) {
		mockUserService := &mocks.UserService{}

		userController := controllers.NewUserController(mockUserService)

		g := SetUpRouterUser(userController)

		req, _ := http.NewRequest(http.MethodPost, "/login", http.NoBody)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		expected := dtos.ResponseMessageOnly{
			Message: constants.RequestBodyInvalid,
		}
		expectedToJson, _ := json.Marshal(expected)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(expectedToJson), rec.Body.String())
		assert.Equal(t, expected.Message, response["message"])
	})

	t.Run("should failed register a new book when user send the same title book", func(t *testing.T) {
		mockUserService := &mocks.UserService{}

		var (
			emailBody    = "frieren@gmail.com"
			passwordBody = "12345"
		)

		newUser := dtos.RequestLoginUser{
			Email:    emailBody,
			Password: passwordBody,
		}

		mockUserService.On("PostLoginUserService", mock.AnythingOfType("*gin.Context"), newUser).Return(nil, apperrors.ErrUserInvalidEmailPassword)
		userController := controllers.NewUserController(mockUserService)

		g := SetUpRouterUser(userController)
		g.Use(middlewares.ErrorHandler)

		body, _ := json.Marshal(newUser)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		expected := dtos.ResponseMessageOnly{
			Message: constants.UserInvalidEmailPassword,
		}
		expectedToJson, _ := json.Marshal(expected)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(expectedToJson), rec.Body.String())
	})
}
