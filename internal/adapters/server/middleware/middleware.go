package middleware

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/authorization"
	"github.com/Inspirate789/Thermy-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ErrorHandler(log logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			for _, ginErr := range ctx.Errors {
				log.Print(logger.LogRecord{
					Name: "Middleware",
					Type: logger.Error,
					Msg:  ginErr.Err.Error(),
				})
			}

			// Put the last error message (possible fatal) to response body
			ctx.JSON(-1, gin.H{"error": ctx.Errors[len(ctx.Errors)-1].Err.Error()}) // -1 not overwrite HTTP status
		}
	}
}

func RoleCheck(svc *authorization.AuthorizationService, parseRole func(*gin.Context) (string, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		requiredRole, err := parseRole(ctx)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		sessionRole, err := svc.GetSessionRole(token)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
		if requiredRole != sessionRole {
			ctx.AbortWithError(http.StatusBadRequest, ErrInvalidRole(sessionRole, requiredRole))
		}

		ctx.Next()
	}
}
