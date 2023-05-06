package server

import (
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// postUnits godoc
//
//	@Summary		Add new units in the given text markup layer.
//	@Description	add new units in the given text markup layer
//	@Tags			student
//	@Param			token		query	string					true	"User authentication token"
//	@Param			layer		query	string					true	"Text markup layer"
//	@Param			unitsDTO	body	interfaces.SaveUnitsDTO	true	"Information about stored units"
//	@Accept			json
//	@Success		200
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/units [post]
//	@Router			/educator/units [post]
//	@Router			/student/units [post]
func (s *Server) postUnits(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
		return
	}
	layer := ctx.Query("layer")

	var unitsDTO interfaces.SaveUnitsDTO
	err = ctx.BindJSON(&unitsDTO)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("SaveUnitsDTO"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = s.storageService.SaveUnits(conn, layer, unitsDTO)
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
//	@Tags			student
//	@Param			token		query	string						true	"User authentication token"
//	@Param			layer		query	string						true	"Text markup layer"
//	@Param			unitsDTO	body	interfaces.UpdateUnitsDTO	true	"Information about updated units"
//	@Accept			json
//	@Success		200
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/units [patch]
//	@Router			/educator/units [patch]
//	@Router			/student/units [patch]
func (s *Server) patchUnits(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
		return
	}
	layer := ctx.Query("layer")

	var unitsDTO interfaces.UpdateUnitsDTO
	err = ctx.BindJSON(&unitsDTO)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("UpdateUnitsDTO"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = s.storageService.UpdateUnits(conn, layer, unitsDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// getAllUnits godoc
//
//	@Summary		Show all units in the given text markup layer.
//	@Description	return all units in the given text markup layer
//	@Tags			student
//	@Param			token	query	string	true	"User authentication token"
//	@Param			layer	query	string	true	"Text markup layer"
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputUnitsDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/units/all [put]
//	@Router			/educator/units/all [put]
//	@Router			/student/units/all [put]
func (s *Server) getAllUnits(ctx *gin.Context) {
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

	unitsDTO, err := s.storageService.GetAllUnits(conn, layer)
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
//	@Tags			student
//	@Param			token			query	string					true	"User authentication token"
//	@Param			layer			query	string					true	"Text markup layer"
//	@Param			propertiesID	body	interfaces.ModelsIdDTO	true	"Models ID according to which the search will be performed"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputUnitsDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/units/models [put]
//	@Router			/educator/units/models [put]
//	@Router			/student/units/models [put]
func (s *Server) getUnitsByModels(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
		return
	}
	layer := ctx.Query("layer")

	var modelsID interfaces.ModelsIdDTO
	err = ctx.BindJSON(&modelsID)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("ModelsIdDTO"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	unitsDTO, err := s.storageService.GetUnitsByModels(conn, layer, modelsID)
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
//	@Tags			student
//	@Param			token			query	string						true	"User authentication token"
//	@Param			layer			query	string						true	"Text markup layer"
//	@Param			propertiesID	body	interfaces.PropertiesIdDTO	true	"Properties ID according to which the search will be performed"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputUnitsDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/units/properties [put]
//	@Router			/educator/units/properties [put]
//	@Router			/student/units/properties [put]
func (s *Server) getUnitsByProperties(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
		return
	}
	layer := ctx.Query("layer")

	var propertiesID interfaces.PropertiesIdDTO
	err = ctx.BindJSON(&propertiesID)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("PropertiesIdDTO"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	unitsDTO, err := s.storageService.GetUnitsByProperties(conn, layer, propertiesID)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, unitsDTO)
}

// getModels godoc
//
//	@Summary		Show all structural models in the given text markup layer.
//	@Description	return all structural models in the given text markup layer
//	@Tags			student
//	@Param			token	query	string	true	"User authentication token"
//	@Param			layer	query	string	true	"Text markup layer"
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputModelsDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/models/all [get]
//	@Router			/educator/models/all [get]
//	@Router			/student/models/all [get]
func (s *Server) getModels(ctx *gin.Context) {
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

	modelsDTO, err := s.storageService.GetModels(conn, layer)
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
//	@Tags			student
//	@Param			token	query	string	true	"User authentication token"
//	@Param			layer	query	string	true	"Text markup layer"
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputModelsDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/elements/all [get]
//	@Router			/educator/elements/all [get]
//	@Router			/student/elements/all [get]
func (s *Server) getModelElements(ctx *gin.Context) {
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

	modelElementsDTO, err := s.storageService.GetModelElements(conn, layer)
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
//	@Tags			student
//	@Param			token	query	string	true	"User authentication token"
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputModelsDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/properties/all [put]
//	@Router			/educator/properties/all [put]
//	@Router			/student/properties/all [put]
func (s *Server) getProperties(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	propertiesDTO, err := s.storageService.GetProperties(conn)
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
//	@Tags			student
//	@Param			token			query	string						true	"User authentication token"
//	@Param			layer			query	string						true	"Text markup layer"
//	@Param			propertiesID	body	interfaces.SearchUnitDTO	true	"Unit data according to which the search will be performed"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.OutputPropertiesDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/properties/unit [put]
//	@Router			/educator/properties/unit [put]
//	@Router			/student/properties/unit [put]
func (s *Server) getPropertiesByUnit(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
		return
	}
	layer := ctx.Query("layer")

	var unitDTO interfaces.SearchUnitDTO
	err = ctx.BindJSON(&unitDTO)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("SearchUnitDTO"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	propertiesDTO, err := s.storageService.GetPropertiesByUnit(conn, layer, unitDTO)
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
//	@Tags			student
//	@Param			token			query	string						true	"User authentication token"
//	@Param			propertyNames	body	interfaces.PropertyNamesDTO	true	"Unit property names"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interfaces.PropertiesIdDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/properties [post]
//	@Router			/educator/properties [post]
//	@Router			/student/properties [post]
func (s *Server) postProperties(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
	}

	var propertyNames interfaces.PropertyNamesDTO
	err = ctx.BindJSON(&propertyNames)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("PropertyNamesDTO"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	propertiesID, err := s.storageService.SaveProperties(conn, propertyNames)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, propertiesID)
}

// getAllLayers godoc
//
//	@Summary		Show all text markup layers.
//	@Description	return all text markup layers
//	@Tags			student
//	@Param			token	query	string	true	"User authentication token"
//	@Produce		json
//	@Success		200	{object}	interfaces.LayersDTO
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/admin/layers/all [get]
//	@Router			/educator/layers/all [get]
//	@Router			/student/layers/all [get]
func (s *Server) getAllLayers(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	layers, err := s.storageService.GetLayers(conn)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, layers)
}
