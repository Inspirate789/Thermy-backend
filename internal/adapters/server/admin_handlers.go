package server

import (
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/pkg/monitoring"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// postUser godoc
//
//	@Summary		Add new user.
//	@Description	add new user
//	@Tags			admin
//	@Param			token	query	string				true	"User authentication token"
//	@Param			user	body	interfaces.UserDTO	true	"User information"
//	@Accept			json
//	@Success		200
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/users [post]
func (s *Server) postUser(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseURLWrap("token"))
		return
	}

	var user interfaces.UserDTO
	err = ctx.BindJSON(&user)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("UserDTO"))
		return
	}

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = s.storageService.AddUser(conn, user)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// getStat godoc
//
//	@Summary		Show the status of server.
//	@Description	return the statistic of the server process
//	@Tags			admin
//	@Param			token	query	string	true	"User authentication token"
//	@Produce		json
//	@Success		200	{object}	monitoring.ProcStat
//	@Failure		500	{object}	string
//	@Router			/admin/stat [get]
func (s *Server) getStat(ctx *gin.Context) {
	observer, err := monitoring.NewProcStatObserver()
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, errors.ErrAccessSystemInfo)
		return
	}

	stat, err := observer.GetInfo()
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, errors.ErrAccessSystemInfo)
		return
	}

	ctx.JSON(http.StatusOK, stat)
}
