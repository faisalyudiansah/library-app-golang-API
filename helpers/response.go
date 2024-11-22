package helpers

import (
	"reflect"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/dtos"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PrintError(c *gin.Context, statusCode int, msg string) {
	res := dtos.ResponseMessageOnly{
		Message: msg,
	}
	c.AbortWithStatusJSON(statusCode, res)
}

func PrintResponse(c *gin.Context, statusCode int, res interface{}) {
	c.JSON(statusCode, res)
}

func FormatterManyBook(books []dtos.ResponseBook, msg string) dtos.ResponseManyData {
	mapBook := map[string][]dtos.ResponseBook{}
	if len(books) == 0 {
		books = []dtos.ResponseBook{}
	}
	mapBook["books"] = books
	res := dtos.ResponseManyData{
		Message:   msg,
		TotalData: int64(len(books)),
		Data:      mapBook,
	}
	return res
}

func FormatterOneBorrow(data *dtos.ResponseBorrow, msg string) dtos.ResponseOneDataBorrow {
	mapBook := map[string]dtos.ResponseBorrow{}
	mapBook["data"] = *data
	res := dtos.ResponseOneDataBorrow{
		Message: msg,
		Result:  mapBook,
	}
	return res
}

func FormatterSuccessRegisterLogin(data *dtos.ResponseDataUser, msg string) dtos.ResponseRegisterUser {
	mapBook := map[string]dtos.ResponseDataUser{}
	mapBook["data"] = *data
	res := dtos.ResponseRegisterUser{
		Message: msg,
		Result:  mapBook,
	}
	return res
}

func FormatterOneBook(book dtos.ResponseBook, msg string) dtos.ResponseOneData {
	mapBook := map[string]dtos.ResponseBook{}
	mapBook["books"] = book
	res := dtos.ResponseOneData{
		Message:   msg,
		TotalData: 1,
		Data:      mapBook,
	}
	return res
}

func FormatterErrorInput(ve validator.ValidationErrors) []dtos.ResponseApiError {
	result := make([]dtos.ResponseApiError, len(ve))
	for i, fe := range ve {
		result[i] = dtos.ResponseApiError{
			Field: jsonFieldName(fe.Field()),
			Msg:   msgForTag(fe.Tag()),
		}
	}
	return result
}

func jsonFieldName(fieldName string) string {
	t := reflect.TypeOf(dtos.RequestValidationMiddleware{})
	field, found := t.FieldByName(fieldName)
	if !found {
		return ""
	}
	jsonTag := field.Tag.Get("json")
	return jsonTag
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return ""
}
