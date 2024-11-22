package servers

import (
	apperrors "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/appErrors"
	"github.com/gin-gonic/gin"
)

func InvalidRoute(c *gin.Context) {
	c.Error(apperrors.ErrUrlNotFound)
}
