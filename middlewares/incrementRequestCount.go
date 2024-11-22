package middlewares

import (
	"strconv"

	"library-api/middlewares/metrics"

	"github.com/gin-gonic/gin"
)

func IncrementRequestCount(c *gin.Context) {
	c.Next()
	metrics.RequestCount.WithLabelValues(c.Request.Method, c.FullPath(), strconv.Itoa(c.Writer.Status())).Inc()
}
