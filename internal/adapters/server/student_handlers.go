package server

import (
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

// postUnits godoc
//
//	@Summary		Add new units in the given text markup layer.
//	@Description	add new units in the given text markup layer
//	@Tags			Units
//	@Param			token		header	string					true	"User authentication token"
//	@Param			layer		query	string					true	"Text markup layer"
//	@Param			unitsDTO	body	interfaces.SaveUnitsDTO	true	"Information about stored units"
//	@Accept			json
//	@Success		200
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/units [post]
func (s *Server) postUnits(ctx *gin.Context) {
	layer := ctx.Query("layer")

	var unitsDTO interfaces.SaveUnitsDTO
	err := ctx.BindJSON(&unitsDTO)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("SaveUnitsDTO"))
		return
	}

	err = s.storageService.SaveUnits(layer, unitsDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// patchUnits godoc
//
//	@Summary		Update existing units in the given text markup layer.
//	@Description	update existing units in the given text markup layer
//	@Tags			Units
//	@Param			token		header	string						true	"User authentication token"
//	@Param			layer		query	string						true	"Text markup layer"
//	@Param			unitsDTO	body	interfaces.UpdateUnitsDTO	true	"Information about updated units"
//	@Accept			json
//	@Success		200
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/units [patch]
func (s *Server) patchUnits(ctx *gin.Context) {
	layer := ctx.Query("layer")

	var unitsDTO interfaces.UpdateUnitsDTO
	err := ctx.BindJSON(&unitsDTO)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("UpdateUnitsDTO"))
		return
	}

	err = s.storageService.UpdateUnits(layer, unitsDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// getUnits godoc
//
//	@Summary		Show all units in the given text markup layer.
//	@Description	return all units in the given text markup layer
//	@Tags			Units
//	@Param			token	header	string	true	"User authentication token"
//	@Param			layer	query	string	true	"Text markup layer"
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputUnitsDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/units [get]
func (s *Server) getUnits(ctx *gin.Context) {
	layer := ctx.Query("layer")

	unitsDTO, err := s.storageService.GetAllUnits(layer)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, unitsDTO)
}

// getUnitsByModels godoc
//
//	@Summary		Show all units with given structural models in the given text markup layer.
//	@Description	return all units with given structural models in the given text markup layer
//	@Tags			Units
//	@Param			token			header	string					true	"User authentication token"
//	@Param			layer			query	string					true	"Text markup layer"
//	@Param			propertiesID	body	interfaces.ModelsIdDTO	true	"Models ID according to which the search will be performed"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputUnitsDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/units/models [put]
func (s *Server) getUnitsByModels(ctx *gin.Context) {
	layer := ctx.Query("layer")

	var modelsID interfaces.ModelsIdDTO
	err := ctx.BindJSON(&modelsID)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("ModelsIdDTO"))
		return
	}

	unitsDTO, err := s.storageService.GetUnitsByModels(layer, modelsID)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, unitsDTO)
}

// getUnitsByProperties godoc
//
//	@Summary		Show all units with given properties in the given text markup layer.
//	@Description	return all units with given properties in the given text markup layer
//	@Tags			Units
//	@Param			token			header	string						true	"User authentication token"
//	@Param			layer			query	string						true	"Text markup layer"
//	@Param			propertiesID	body	interfaces.PropertiesIdDTO	true	"Properties ID according to which the search will be performed"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputUnitsDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/units/properties [put]
func (s *Server) getUnitsByProperties(ctx *gin.Context) {
	layer := ctx.Query("layer")

	var propertiesID interfaces.PropertiesIdDTO
	err := ctx.BindJSON(&propertiesID)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("PropertiesIdDTO"))
		return
	}

	unitsDTO, err := s.storageService.GetUnitsByProperties(layer, propertiesID)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, unitsDTO)
}

// deleteUnits godoc
//
//	@Summary		Delete existing units in the given text markup layer.
//	@Description	delete existing units in the given text markup layer
//	@Tags			Units
//	@Param			token		header	string						true	"User authentication token"
//	@Param			layer		query	string						true	"Text markup layer"
//	@Param			unitsDTO	body	interfaces.SearchUnitDTO	true	"Information about updated units"
//	@Accept			json
//	@Success		200
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/units [delete]
func (s *Server) deleteUnits(ctx *gin.Context) { // TODO
	layer := ctx.Query("layer")

	var unitDTO interfaces.SearchUnitDTO
	err := ctx.BindJSON(&unitDTO)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("SearchUnitDTO"))
		return
	}

	propertiesDTO, err := s.storageService.GetPropertiesByUnit(layer, unitDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, propertiesDTO)
}

// getModels godoc
//
//	@Summary		Show all structural models in the given text markup layer.
//	@Description	return all structural models in the given text markup layer
//	@Tags			Models
//	@Param			token	header	string	true	"User authentication token"
//	@Param			layer	query	string	true	"Text markup layer"
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputModelsDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/models [get]
func (s *Server) getModels(ctx *gin.Context) {
	layer := ctx.Query("layer")

	modelsDTO, err := s.storageService.GetModels(layer)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, modelsDTO)
}

// getModelElements godoc
//
//	@Summary		Show all elements of structural models in the given text markup layer.
//	@Description	return all elements of structural models in the given text markup layer
//	@Tags			Elements
//	@Param			token	header	string	true	"User authentication token"
//	@Param			layer	query	string	true	"Text markup layer"
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputModelsDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/elements [get]
func (s *Server) getModelElements(ctx *gin.Context) {
	layer := ctx.Query("layer")

	modelElementsDTO, err := s.storageService.GetModelElements(layer)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, modelElementsDTO)
}

// getProperties godoc
//
//	@Summary		Show all unit properties.
//	@Description	return all unit properties
//	@Tags			Properties
//	@Param			token	header	string	true	"User authentication token"
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputModelsDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/properties [get]
func (s *Server) getProperties(ctx *gin.Context) {

	propertiesDTO, err := s.storageService.GetProperties()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, propertiesDTO)
}

// getPropertiesByUnit godoc
//
//	@Summary		Show all properties for the given unit in the given text markup layer.
//	@Description	return all properties for the given unit in the given text markup layer
//	@Tags			Properties
//	@Param			token			header	string						true	"User authentication token"
//	@Param			layer			query	string						true	"Text markup layer"
//	@Param			propertiesID	body	interfaces.SearchUnitDTO	true	"Unit data according to which the search will be performed"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputPropertiesDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/properties/unit [put]
func (s *Server) getPropertiesByUnit(ctx *gin.Context) {
	layer := ctx.Query("layer")

	var unitDTO interfaces.SearchUnitDTO
	err := ctx.BindJSON(&unitDTO)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("SearchUnitDTO"))
		return
	}

	propertiesDTO, err := s.storageService.GetPropertiesByUnit(layer, unitDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, propertiesDTO)
}

// postProperties godoc
//
//	@Summary		Add new unit properties.
//	@Description	add new unit properties
//	@Tags			Properties
//	@Param			token			header	string						true	"User authentication token"
//	@Param			propertyNames	body	interfaces.PropertyNamesDTO	true	"Unit property names"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.PropertiesIdDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/properties [post]
func (s *Server) postProperties(ctx *gin.Context) {
	var propertyNames interfaces.PropertyNamesDTO
	err := ctx.BindJSON(&propertyNames)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("PropertyNamesDTO"))
		return
	}

	propertiesID, err := s.storageService.SaveProperties(propertyNames)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, propertiesID)
}

// getLayers godoc
//
//	@Summary		Show all text markup layers.
//	@Description	return all text markup layers
//	@Tags			Layers
//	@Param			token	header	string	true	"User authentication token"
//	@Produce		json
//	@Success		200	{object}	interfaces.LayersDTO
//	@Failure		400	{object}	string
//	@Failure		401	{object}	string
//	@Failure		500	{object}	string
//	@Router			/layers [get]
func (s *Server) getLayers(ctx *gin.Context) {
	layers, err := s.storageService.GetLayers()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, layers)
}
