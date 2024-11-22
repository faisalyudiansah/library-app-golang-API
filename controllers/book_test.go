package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	apperrors "library-api/appErrors"
	"library-api/constants"
	"library-api/controllers"
	"library-api/dtos"
	"library-api/helpers"
	"library-api/middlewares"
	"library-api/mocks"
	"library-api/models"
	"library-api/servers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SetUpRouter(bookController *controllers.BookController) *gin.Engine {
	h := &servers.HandlerOps{
		BookController: bookController,
	}
	return servers.SetupRoute(h)
}

func TestGetAllBooksControllers(t *testing.T) {
	t.Run("should success get all books when user hit api get all books", func(t *testing.T) {
		mockBookService := &mocks.BookService{}
		coverInput := "Hardcover"
		books := []models.Book{
			{
				ID:          1,
				AuthorId:    1,
				Title:       "Wehehe",
				Description: "Lapar bang",
				Quantity:    12,
				Cover:       &coverInput,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				DeleteAt:    nil,
			},
			{
				ID:          2,
				AuthorId:    2,
				Title:       "Wihihi",
				Description: "Haus bang",
				Quantity:    2,
				Cover:       &coverInput,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				DeleteAt:    nil,
			},
		}
		booksDto := dtos.ToResponseBookType(books)
		mockBookService.On("GetAllBookService", mock.AnythingOfType("*gin.Context"), "").Return(booksDto, nil)
		bookController := controllers.NewBookController(mockBookService)

		g := SetUpRouter(bookController)

		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		expected := helpers.FormatterManyBook(dtos.ToResponseBookType(books), constants.Ok)
		expectedToJson, _ := json.Marshal(expected)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedToJson), rec.Body.String())
		assert.Equal(t, constants.Ok, response["message"])
		mockBookService.AssertNumberOfCalls(t, "GetAllBookService", 1)
	})

	t.Run("should error get all books when user hit api get all books", func(t *testing.T) {
		mockBookService := &mocks.BookService{}
		mockBookService.On("GetAllBookService", mock.AnythingOfType("*gin.Context"), "").Return(nil, apperrors.ErrISE)
		bookController := controllers.NewBookController(mockBookService)
		g := SetUpRouter(bookController)

		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, constants.ISE, response["message"])
		mockBookService.AssertNumberOfCalls(t, "GetAllBookService", 1)
	})

	t.Run("should not return an error if data book is empty", func(t *testing.T) {
		mockBookService := &mocks.BookService{}
		mockBookService.On("GetAllBookService", mock.AnythingOfType("*gin.Context"), "").Return([]dtos.ResponseBook{}, nil)
		bookController := controllers.NewBookController(mockBookService)

		g := SetUpRouter(bookController)

		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, constants.Ok, response["message"])
		mockBookService.AssertNumberOfCalls(t, "GetAllBookService", 1)
	})
}

func TestPostBooksControllers(t *testing.T) {
	t.Run("should success create a new book when user send correct body and hit api", func(t *testing.T) {
		mockBookService := &mocks.BookService{}

		var (
			titlePost       = "New Book"
			authorId        = int64(1)
			descriptionPost = "Buku baru nihhhh"
			quantityPost    = int64(12)
			coverPost       = "Hardcover"
		)

		newProduct := dtos.RequestPostBook{
			Title:       titlePost,
			AuthorId:    authorId,
			Description: descriptionPost,
			Quantity:    quantityPost,
			Cover:       &coverPost,
		}

		expectedProduct := models.Book{
			ID:          1,
			Title:       titlePost,
			Description: descriptionPost,
			Quantity:    quantityPost,
			Cover:       &coverPost,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeleteAt:    nil,
		}
		expectedProductRes := dtos.ToDtoResponseBook(expectedProduct)
		mockBookService.On("PostBookService", mock.AnythingOfType("*gin.Context"), newProduct).Return(expectedProductRes, nil)
		bookController := controllers.NewBookController(mockBookService)

		g := SetUpRouter(bookController)

		reqBody, _ := json.Marshal(newProduct)
		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		expected := helpers.FormatterOneBook(dtos.ToDtoResponseBook(expectedProduct), constants.SuccessAddNewBook)
		expectedToJson, _ := json.Marshal(expected)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expected.Message, response["message"])
		assert.Equal(t, string(expectedToJson), rec.Body.String())
		mockBookService.AssertNumberOfCalls(t, "PostBookService", 1)
	})

	t.Run("should success create a new book when user doesn't send cover field", func(t *testing.T) {
		mockBookService := &mocks.BookService{}

		var (
			titlePost       = "New Book"
			authorId        = int64(1)
			descriptionPost = "Buku baru nihhhh"
			quantityPost    = int64(12)
		)

		newProduct := dtos.RequestPostBook{
			Title:       titlePost,
			AuthorId:    authorId,
			Description: descriptionPost,
			Quantity:    quantityPost,
		}

		expectedProduct := models.Book{
			ID:          1,
			Title:       titlePost,
			Description: descriptionPost,
			Quantity:    quantityPost,
			Cover:       nil,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeleteAt:    nil,
		}
		expectedProductRes := dtos.ToDtoResponseBook(expectedProduct)
		mockBookService.On("PostBookService", mock.AnythingOfType("*gin.Context"), newProduct).Return(expectedProductRes, nil)
		bookController := controllers.NewBookController(mockBookService)

		g := SetUpRouter(bookController)

		reqBody, _ := json.Marshal(newProduct)
		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		expected := helpers.FormatterOneBook(dtos.ToDtoResponseBook(expectedProduct), constants.SuccessAddNewBook)
		expectedToJson, _ := json.Marshal(expected)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expected.Message, response["message"])
		assert.Equal(t, string(expectedToJson), rec.Body.String())
		mockBookService.AssertNumberOfCalls(t, "PostBookService", 1)
	})

	t.Run("should error when user doesnt send request body", func(t *testing.T) {
		mockBookService := &mocks.BookService{}

		bookController := controllers.NewBookController(mockBookService)

		g := SetUpRouter(bookController)

		req, _ := http.NewRequest(http.MethodPost, "/books", http.NoBody)
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
		assert.Equal(t, expected.Message, response["message"])
		assert.Equal(t, string(expectedToJson), rec.Body.String())
	})

	t.Run("should failed create a new book when user send the same title book", func(t *testing.T) {
		mockBookService := &mocks.BookService{}

		var (
			titlePost       = "Buku yang sudah ada bree"
			authorId        = int64(1)
			descriptionPost = "Buku bro nihhhh"
			quantityPost    = int64(12)
			coverPost       = "Hardcover"
		)

		newProduct := dtos.RequestPostBook{
			Title:       titlePost,
			AuthorId:    authorId,
			Description: descriptionPost,
			Quantity:    quantityPost,
			Cover:       &coverPost,
		}
		mockBookService.On("PostBookService", mock.AnythingOfType("*gin.Context"), newProduct).Return(dtos.ResponseBook{}, apperrors.ErrTitleBookAlreadyExists)
		bookController := controllers.NewBookController(mockBookService)

		g := SetUpRouter(bookController)

		reqBody, _ := json.Marshal(newProduct)
		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		expected := dtos.ResponseMessageOnly{
			Message: constants.FailedTitleAlreadyExists,
		}
		expectedToJson, _ := json.Marshal(expected)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, expected.Message, response["message"])
		assert.Equal(t, string(expectedToJson), rec.Body.String())
		mockBookService.AssertNumberOfCalls(t, "PostBookService", 1)
	})

	t.Run("should fail to create a new book when user does not send title or any mandatory fields", func(t *testing.T) {
		mockBookService := &mocks.BookService{}
		bookController := controllers.NewBookController(mockBookService)
		g := SetUpRouter(bookController)
		g.Use(middlewares.ErrorHandler)

		newProductWithoutTitle := dtos.RequestPostBook{
			AuthorId:    1,
			Description: "Buku bro nihhhh",
			Quantity:    12,
		}

		body, _ := json.Marshal(newProductWithoutTitle)
		req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		g.ServeHTTP(rec, req)

		var response map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &response)

		expected := map[string]interface{}{
			"errors": []dtos.ResponseApiError{
				{
					Field: "title",
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
		assert.Equal(t, "title", errorDetail["field"])
		assert.Equal(t, "This field is required", errorDetail["message"])
	})
}
