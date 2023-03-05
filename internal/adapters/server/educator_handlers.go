package server

import (
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) postModels(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL"))
		return
	}
	layer := ctx.Query("layer")

	var modelNames interfaces.ModelNamesDTO
	err = ctx.BindJSON(&modelNames)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse ModelNamesDTO from received JSON"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	modelsID, err := s.storageService.SaveModels(conn, layer, modelNames)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, modelsID)
}

func (s *Server) postElements(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL"))
		return
	}
	layer := ctx.Query("layer")

	var modelElementNames interfaces.ModelElementNamesDTO
	err = ctx.BindJSON(&modelElementNames)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse ModelElementNamesDTO from received JSON"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	modelElementsID, err := s.storageService.SaveModelElements(conn, layer, modelElementNames)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, modelElementsID)
}

func (s *Server) postLayer(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL"))
		return
	}
	layer := ctx.Query("layer")

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = s.storageService.SaveLayer(conn, layer)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": "ok"})
}
