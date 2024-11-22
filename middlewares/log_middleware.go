package middlewares

import (
	"fmt"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/helpers/logger"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log := logger.Log
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		fmt.Printf("ERRORS: %+v\n", ctx.Errors)

		if lastErr := ctx.Errors.Last(); lastErr != nil {
			log.WithFields(map[string]any{
				"METHOD":    reqMethod,
				"URI":       reqUri,
				"STATUS":    statusCode,
				"LATENCY":   latencyTime,
				"CLIENT_IP": clientIP,
			}).Error(ctx.Errors)
			return
		}

		log.WithFields(map[string]any{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}).Infof("REQUEST %s %s SUCCESS", reqMethod, reqUri)
	}
}
