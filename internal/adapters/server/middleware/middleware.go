package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ErrorResponseWriter(_ *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			// Put the last error message (possible fatal) to response body
			// ctx.JSON(-1, gin.H{"error": ctx.Errors[len(ctx.Errors)-1].Err.Error()}) // -1 not overwrite HTTP status
			ctx.JSON(-1, ctx.Errors[len(ctx.Errors)-1].Err.Error()) // -1 not overwrite HTTP status
		}
	}
}
