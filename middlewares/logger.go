package middlewares

import (
	"zhiHu/logger"
	"time"
	// "fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc{
	return func (ctx *gin.Context) {
		startTime := time.Now().UnixNano() / 1e6
		path := ctx.Request.URL.Path
		method := ctx.Request.Method

		ctx.Next()
		
		status := ctx.Writer.Status()
		now := time.Now().Format("2006-01-02 - 15:04:05")
		endTime := time.Now().UnixNano() / 1e6
		cost := endTime - startTime
		logger.Info("%v | %v | %v | %v | %v", now, status, cost, method, path)
	}
}