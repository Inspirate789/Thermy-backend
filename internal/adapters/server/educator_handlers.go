package server

import (
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// postModels godoc
//
//	@Summary		Add new structural models.
//	@Description	add new structural models
//	@Tags			educator
//	@Param			token		query	string						true	"User authentication token"
//	@Param			layer		query	string						true	"Text markup layer"
//	@Param			modelNames	body	interfaces.ModelNamesDTO	true	"Structural model names"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.ModelsIdDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/models [post]
func (s *Server) postModels(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
		return
	}
	layer := ctx.Query("layer")

	var modelNames interfaces.ModelNamesDTO
	err = ctx.BindJSON(&modelNames)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("ModelNamesDTO"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	modelsID, err := s.storageService.SaveModels(conn, layer, modelNames)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, modelsID)
}

// postElements godoc
//
//	@Summary		Add new elements of structural models.
//	@Description	add new elements of structural models
//	@Tags			educator
//	@Param			token				query	string							true	"User authentication token"
//	@Param			layer				query	string							true	"Text markup layer"
//	@Param			modelElementNames	body	interfaces.ModelElementNamesDTO	true	"Structural model element names"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.ModelElementsIdDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/elements [post]
func (s *Server) postElements(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
		return
	}
	layer := ctx.Query("layer")

	var modelElementNames interfaces.ModelElementNamesDTO
	err = ctx.BindJSON(&modelElementNames)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("ModelElementNamesDTO"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	modelElementsID, err := s.storageService.SaveModelElements(conn, layer, modelElementNames)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, modelElementsID)
}

// postLayer godoc
//
//	@Summary		Add new text markup layer.
//	@Description	add new text markup layer
//	@Tags			educator
//	@Param			token	query	string	true	"User authentication token"
//	@Param			layer	query	string	true	"Text markup layer"
//	@Success		200
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/layers [post]
func (s *Server) postLayer(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
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
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
