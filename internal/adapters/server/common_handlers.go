package server

import (
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) login(ctx *gin.Context) {
	var request entities.AuthRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse AuthRequest from received JSON"))
		return
	}

	token, err := s.authService.AddSession(s.storageService, &request, ctx)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": strconv.FormatUint(token, 10),
	})
}

func (s *Server) logout(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL"))
		return
	}

	err = s.authService.RemoveSession(s.storageService, token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.Status(http.StatusOK)
}
