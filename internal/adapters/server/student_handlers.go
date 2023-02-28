package server

import (
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) postUnits(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL")) // TODO: log true error message, return readable error message (use wrap?)
		return
	}
	layer := ctx.Param("layer")

	var unitsDTO interfaces.SaveUnitsDTO
	err = ctx.BindJSON(&unitsDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse SaveUnitsDTO from received JSON")) // TODO: add custom errors
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = s.storageService.SaveUnits(conn, layer, unitsDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": "ok"})
}

func (s *Server) patchUnits(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL")) // TODO: log true error message, return readable error message (use wrap?)
		return
	}
	layer := ctx.Param("layer")

	var unitsDTO interfaces.UpdateUnitsDTO
	err = ctx.BindJSON(&unitsDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse UpdateUnitsDTO from received JSON")) // TODO: add custom errors
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = s.storageService.UpdateUnits(conn, layer, unitsDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": "ok"})
}

func (s *Server) getAllUnits(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL")) // TODO: log true error message, return readable error message (use wrap)
		return
	}
	layer := ctx.Param("layer")

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	unitsDTO, err := s.storageService.GetAllUnits(conn, layer)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, unitsDTO)
}

func (s *Server) getUnitsByModels(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL")) // TODO: log true error message, return readable error message (use wrap)
		return
	}
	layer := ctx.Param("layer")

	var modelsID interfaces.ModelsIdDTO
	err = ctx.BindJSON(&modelsID)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse ModelsIdDTO from received JSON"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	unitsDTO, err := s.storageService.GetUnitsByModels(conn, layer, modelsID)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, unitsDTO)
}

func (s *Server) getUnitsByProperties(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL")) // TODO: log true error message, return readable error message (use wrap)
		return
	}
	layer := ctx.Param("layer")

	var propertiesID interfaces.PropertiesIdDTO
	err = ctx.BindJSON(&propertiesID)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse PropertiesIdDTO from received JSON"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	unitsDTO, err := s.storageService.GetUnitsByProperties(conn, layer, propertiesID)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, unitsDTO)
}

func (s *Server) getModels(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL")) // TODO: log true error message, return readable error message (use wrap)
		return
	}
	layer := ctx.Param("layer")

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	modelsDTO, err := s.storageService.GetModels(conn, layer)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, modelsDTO)
}

func (s *Server) getModelElements(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL")) // TODO: log true error message, return readable error message (use wrap)
		return
	}
	layer := ctx.Param("layer")

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	modelElementsDTO, err := s.storageService.GetModelElements(conn, layer)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, modelElementsDTO)
}

func (s *Server) getProperties(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL")) // TODO: log true error message, return readable error message (use wrap)
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	propertiesDTO, err := s.storageService.GetProperties(conn)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, propertiesDTO)
}

func (s *Server) getPropertiesByUnit(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL"))
		return
	}
	layer := ctx.Param("layer")

	var unitDTO interfaces.SearchUnitDTO
	err = ctx.BindJSON(&unitDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse SearchUnitDTO from received JSON"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	propertiesDTO, err := s.storageService.GetPropertiesByUnit(conn, layer, unitDTO)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, propertiesDTO)
}

func (s *Server) postProperties(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL")) // TODO: log true error message, return readable error message (use wrap)
		return
	}

	var propertyNames interfaces.PropertyNamesDTO
	err = ctx.BindJSON(&propertyNames)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse PropertyNamesDTO from received JSON"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	propertiesID, err := s.storageService.SaveProperties(conn, propertyNames)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, propertiesID)
}

func (s *Server) getAllLayers(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Param("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL")) // TODO: log true error message, return readable error message (use wrap)
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	layers, err := s.storageService.GetLayers(conn)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, layers)
}