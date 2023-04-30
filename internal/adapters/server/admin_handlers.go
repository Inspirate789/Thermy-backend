package server

import (
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/pkg/monitoring"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.Status(http.StatusOK)
}

//func (s *Server) getUserPassword(ctx *gin.Context) {
//	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
//	if err != nil {
//		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL"))
//		return
//	}
//	username := ctx.Query("username")
//
//	conn, err := s.authService.GetSessionConn(token)
//	if err != nil {
//		_ = ctx.AbortWithError(http.StatusBadRequest, err)
//		return
//	}
//
//	password, err := s.storageService.GetUserPassword(conn, username)
//	if err != nil {
//		_ = ctx.AbortWithError(http.StatusBadRequest, err)
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"password": password, "error": "ok"})
//}

func (s *Server) getStat(ctx *gin.Context) {
	observer, err := monitoring.NewProcStatObserver()
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrAccessSystemInfo)
		return
	}

	stat, err := observer.GetInfo()
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrAccessSystemInfo)
		return
	}

	ctx.JSON(http.StatusOK, stat)
}
