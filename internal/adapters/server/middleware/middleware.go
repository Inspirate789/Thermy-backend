package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

func RoleCheck(requiredRole string) gin.HandlerFunc { // TODO
	return func(ctx *gin.Context) {
		token, err := strconv.ParseUint(ctx.GetHeader("token"), 10, 64)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		sessionRole, err := svc.GetSessionRole(token)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if requiredRole != sessionRole {
			_ = ctx.AbortWithError(http.StatusBadRequest, ErrInvalidRole(requiredRole, sessionRole))
			return
		}

		ctx.Next()
	}
}
