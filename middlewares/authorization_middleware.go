package middlewares

import (
	"context"
	"strings"

	apperrors "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/appErrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/models"

	// "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/models"
	"github.com/gin-gonic/gin"
)

func AuthorizationBorrow(c *gin.Context) {
	reqToken := c.GetHeader("Authorization")
	if reqToken == "" || len(reqToken) == 0 {
		c.Error(apperrors.ErrUnauthorization)
		c.Abort()
		return
	}
	splitToken := strings.Split(reqToken, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		c.Error(apperrors.ErrUnauthorization)
		c.Abort()
		return
	}
	jwtProvider := helpers.NewJWTProviderHS256()
	result, err := jwtProvider.VerifyToken(splitToken[1])
	if err != nil {
		c.Error(apperrors.ErrUnauthorization)
		c.Abort()
		return
	}

	var userId models.ID = "userId"
	ctx := context.WithValue(c.Request.Context(), userId, result.UserID)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
