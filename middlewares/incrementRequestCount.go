package middlewares

import (
	"strconv"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/middlewares/metrics"
	"github.com/gin-gonic/gin"
)

func IncrementRequestCount(c *gin.Context) {
	c.Next()
	metrics.RequestCount.WithLabelValues(c.Request.Method, c.FullPath(), strconv.Itoa(c.Writer.Status())).Inc()
}
