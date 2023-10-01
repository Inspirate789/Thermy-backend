package server

import (
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

// postModels godoc
//
//	@Summary		Add new structural models.
//	@Description	add new structural models
//	@Tags			Models
//	@Security		ApiKeyAuth
//	@Param			Authorization	header	string						true	"Authorization"
//	@Param			layer			query	string						true	"Text markup layer"
//	@Param			modelNames		body	interfaces.ModelNamesDTO	true	"Structural model names"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.ModelsIdDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/models [post]
func (s *Server) postModels(ctx *gin.Context) {
	layer := ctx.Query("layer")

	var modelNames interfaces.ModelNamesDTO
	err := ctx.BindJSON(&modelNames)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("ModelNamesDTO"))
		return
	}

	modelsID, err := s.storageService.SaveModels(layer, modelNames)
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
//	@Tags			Elements
//	@Security		ApiKeyAuth
//	@Param			Authorization		header	string							true	"Authorization"
//	@Param			layer				query	string							true	"Text markup layer"
//	@Param			modelElementNames	body	interfaces.ModelElementNamesDTO	true	"Structural model element names"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.ModelElementsIdDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/elements [post]
func (s *Server) postElements(ctx *gin.Context) {
	layer := ctx.Query("layer")

	var modelElementNames interfaces.ModelElementNamesDTO
	err := ctx.BindJSON(&modelElementNames)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("ModelElementNamesDTO"))
		return
	}

	modelElementsID, err := s.storageService.SaveModelElements(layer, modelElementNames)
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
//	@Tags			Layers
//	@Security		ApiKeyAuth
//	@Param			Authorization	header	string	true	"Authorization"
//	@Param			layer			query	string	true	"Text markup layer"
//	@Success		200
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/layers [post]
func (s *Server) postLayer(ctx *gin.Context) {
	layer := ctx.Query("layer")

	err := s.storageService.SaveLayer(layer)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}
