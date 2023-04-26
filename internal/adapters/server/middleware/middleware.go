package middleware

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/authorization"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func ErrorHandler(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			for _, ginErr := range ctx.Errors {
				log.Error(ginErr.Err.Error())
			}

			// Put the last error message (possible fatal) to response body
			ctx.JSON(-1, gin.H{"error": ctx.Errors[len(ctx.Errors)-1].Err.Error()}) // -1 not overwrite HTTP status
		}
	}
}

func SessionCheck(svc authorization.AuthManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
		}

		if !svc.SessionExist(token) {
			_ = ctx.AbortWithError(http.StatusBadRequest, ErrUserNotExist(ctx.Query("token")))
		}

		ctx.Next()
	}
}

func RoleCheck(svc authorization.AuthManager, parseRole func(*gin.Context) (string, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
		}

		requiredRole, err := parseRole(ctx)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
		}

		sessionRole, err := svc.GetSessionRole(token)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
		}
		if requiredRole != sessionRole {
			_ = ctx.AbortWithError(http.StatusBadRequest, ErrInvalidRole(sessionRole, requiredRole))
		}

		ctx.Next()
	}
}
