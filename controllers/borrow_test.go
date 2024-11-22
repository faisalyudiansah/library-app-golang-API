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
	utilscontext "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers/utilsContext"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/middlewares"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	userId   = int64(1)
	bookId   = int64(2)
	borrowId = int64(3)
)

func TestBorrowController_PostNewBorrowBookController(t *testing.T) {
	tests := []struct {
		name        string
		msg         string
		mockService func(*gin.Context, *mocks.BorrowService, dtos.RequestBorrowBook, *dtos.ResponseBorrow) *controllers.BorrowController
		callMock    bool
		reqBody     interface{}
		resService  *dtos.ResponseBorrow
		wantRes     func(*dtos.ResponseBorrow, string) interface{}
		statusCode  int
	}{
		{
			name: "should be successful when post of making a new book loan",
			msg:  constants.SuccessBorrowBook,
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestBorrowBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				bs.On("PostNewBorrowService", ctx, reqBody, utilscontext.GetValueUserIdFromToken(ctx)).Return(resService, nil)
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock: true,
			reqBody: dtos.RequestBorrowBook{
				BookID:     bookId,
				BorrowDate: &time.Time{},
			},
			resService: &dtos.ResponseBorrow{
				ID:         int64(1),
				UserID:     userId,
				BookID:     bookId,
				BorrowDate: time.Time{},
				ReturnDate: nil,
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
				DeleteAt:   nil,
			},
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				res := helpers.FormatterOneBorrow(rb, s)
				return res
			},
			statusCode: http.StatusCreated,
		},
		{
			name: "should be error when user does not sent any request body",
			msg:  constants.RequestBodyInvalid,
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestBorrowBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock:   false,
			reqBody:    nil,
			resService: nil,
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				return dtos.ResponseMessageOnly{
					Message: s,
				}
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "should be error when user does not sent a mandatory field in request body",
			msg:  "",
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestBorrowBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock: false,
			reqBody: dtos.RequestBorrowBook{
				BorrowDate: &time.Time{},
			},
			resService: nil,
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				res := map[string]interface{}{
					"errors": []dtos.ResponseApiError{
						{
							Field: "book_id",
							Msg:   "This field is required",
						},
					},
				}
				return res
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "should be error when book id is invalid",
			msg:  constants.InvalidBookId,
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestBorrowBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				bs.On("PostNewBorrowService", ctx, reqBody, utilscontext.GetValueUserIdFromToken(ctx)).Return(nil, apperrors.ErrInvalidBookId)
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock: true,
			reqBody: dtos.RequestBorrowBook{
				BookID:     bookId,
				BorrowDate: &time.Time{},
			},
			resService: nil,
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				return dtos.ResponseMessageOnly{
					Message: s,
				}
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "should be error when user set invalid access token / unauthorized",
			msg:  constants.InvalidAccessToken,
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestBorrowBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				bs.On("PostNewBorrowService", ctx, reqBody, utilscontext.GetValueUserIdFromToken(ctx)).Return(nil, apperrors.ErrInvalidAccessToken)
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock: true,
			reqBody: dtos.RequestBorrowBook{
				BookID:     bookId,
				BorrowDate: &time.Time{},
			},
			resService: nil,
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				return dtos.ResponseMessageOnly{
					Message: s,
				}
			},
			statusCode: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Set("userId", userId)

			mockBorrowService := &mocks.BorrowService{}
			var borrowController *controllers.BorrowController
			if tt.reqBody != nil {
				borrowController = tt.mockService(ctx, mockBorrowService, tt.reqBody.(dtos.RequestBorrowBook), tt.resService)
				reqBody, _ := json.Marshal(tt.reqBody)
				req := httptest.NewRequest(http.MethodPost, "/borrow-books", bytes.NewBuffer(reqBody))
				ctx.Request = req
			} else {
				borrowController = tt.mockService(ctx, mockBorrowService, dtos.RequestBorrowBook{}, tt.resService)
				req := httptest.NewRequest(http.MethodPost, "/borrow-books", nil)
				ctx.Request = req
			}

			borrowController.PostNewBorrowBookController(ctx)
			middlewares.ErrorHandler(ctx)

			expectedToJson, err := json.Marshal(tt.wantRes(tt.resService, tt.msg))
			assert.Nil(t, err)
			assert.Equal(t, tt.statusCode, rec.Code)
			assert.Equal(t, string(expectedToJson), rec.Body.String())
			if tt.callMock {
				mockBorrowService.AssertNumberOfCalls(t, "PostNewBorrowService", 1)
			}
		})
	}
}

func TestBorrowController_PostReturnBookController(t *testing.T) {
	tests := []struct {
		name        string
		msg         string
		mockService func(*gin.Context, *mocks.BorrowService, dtos.RequestReturnBook, *dtos.ResponseBorrow) *controllers.BorrowController
		callMock    bool
		reqBody     interface{}
		resService  *dtos.ResponseBorrow
		wantRes     func(*dtos.ResponseBorrow, string) interface{}
		statusCode  int
	}{
		{
			name: "should be successfully when user want to return a book",
			msg:  constants.SuccessReturnBook,
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestReturnBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				bs.On("PostReturnBookService", ctx, reqBody, utilscontext.GetValueUserIdFromToken(ctx)).Return(resService, nil)
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock: true,
			reqBody: dtos.RequestReturnBook{
				BorrowID:   borrowId,
				BookID:     bookId,
				ReturnDate: &time.Time{},
			},
			resService: &dtos.ResponseBorrow{
				ID:         int64(1),
				UserID:     userId,
				BookID:     bookId,
				BorrowDate: time.Time{},
				ReturnDate: &time.Time{},
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
				DeleteAt:   nil,
			},
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				res := helpers.FormatterOneBorrow(rb, s)
				return res
			},
			statusCode: http.StatusOK,
		},
		{
			name: "should be error when user does not sent any request body",
			msg:  constants.RequestBodyInvalid,
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestReturnBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock:   false,
			reqBody:    nil,
			resService: nil,
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				return dtos.ResponseMessageOnly{
					Message: s,
				}
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "should be error when user does not sent a mandatory field in request body",
			msg:  "",
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestReturnBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock: false,
			reqBody: dtos.RequestReturnBook{
				BorrowID:   borrowId,
				ReturnDate: &time.Time{},
			},
			resService: nil,
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				res := map[string]interface{}{
					"errors": []dtos.ResponseApiError{
						{
							Field: "book_id",
							Msg:   "This field is required",
						},
					},
				}
				return res
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "should be error when book id is invalid",
			msg:  constants.InvalidBookId,
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestReturnBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				bs.On("PostReturnBookService", ctx, reqBody, utilscontext.GetValueUserIdFromToken(ctx)).Return(nil, apperrors.ErrInvalidBookId)
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock: true,
			reqBody: dtos.RequestReturnBook{
				BorrowID:   borrowId,
				BookID:     bookId,
				ReturnDate: &time.Time{},
			},
			resService: nil,
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				return dtos.ResponseMessageOnly{
					Message: s,
				}
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "should be error when borrow id is invalid",
			msg:  constants.InvalidBorrowIdBookId,
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestReturnBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				bs.On("PostReturnBookService", ctx, reqBody, utilscontext.GetValueUserIdFromToken(ctx)).Return(nil, apperrors.ErrInvalidBorrowIdBookId)
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock: true,
			reqBody: dtos.RequestReturnBook{
				BorrowID:   borrowId,
				BookID:     bookId,
				ReturnDate: &time.Time{},
			},
			resService: nil,
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				return dtos.ResponseMessageOnly{
					Message: s,
				}
			},
			statusCode: http.StatusBadRequest,
		},
		{
			name: "should be error when user set invalid access token / unauthorized",
			msg:  constants.InvalidAccessToken,
			mockService: func(
				ctx *gin.Context,
				bs *mocks.BorrowService,
				reqBody dtos.RequestReturnBook,
				resService *dtos.ResponseBorrow,
			) *controllers.BorrowController {
				bs.On("PostReturnBookService", ctx, reqBody, utilscontext.GetValueUserIdFromToken(ctx)).Return(nil, apperrors.ErrInvalidAccessToken)
				borrowController := controllers.NewBorrowController(bs)
				return borrowController
			},
			callMock: true,
			reqBody: dtos.RequestReturnBook{
				BorrowID:   borrowId,
				BookID:     bookId,
				ReturnDate: &time.Time{},
			},
			resService: nil,
			wantRes: func(rb *dtos.ResponseBorrow, s string) interface{} {
				return dtos.ResponseMessageOnly{
					Message: s,
				}
			},
			statusCode: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Set("userId", userId)

			mockBorrowService := &mocks.BorrowService{}
			var borrowController *controllers.BorrowController
			if tt.reqBody != nil {
				borrowController = tt.mockService(ctx, mockBorrowService, tt.reqBody.(dtos.RequestReturnBook), tt.resService)
				reqBody, _ := json.Marshal(tt.reqBody)
				req := httptest.NewRequest(http.MethodPost, "/return-books", bytes.NewBuffer(reqBody))
				ctx.Request = req
			} else {
				borrowController = tt.mockService(ctx, mockBorrowService, dtos.RequestReturnBook{}, tt.resService)
				req := httptest.NewRequest(http.MethodPost, "/return-books", nil)
				ctx.Request = req
			}

			borrowController.PostReturnBookController(ctx)
			middlewares.ErrorHandler(ctx)

			expectedToJson, err := json.Marshal(tt.wantRes(tt.resService, tt.msg))
			assert.Nil(t, err)
			assert.Equal(t, tt.statusCode, rec.Code)
			assert.Equal(t, string(expectedToJson), rec.Body.String())
			if tt.callMock {
				mockBorrowService.AssertNumberOfCalls(t, "PostReturnBookService", 1)
			}
		})
	}
}
