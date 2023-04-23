package server

import (
	"errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/pkg/monitoring"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) postUser(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL"))
		return
	}

	var user interfaces.UserDTO
	err = ctx.BindJSON(&user)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse UserDTO from received JSON"))
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
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	stat, err := observer.GetInfo()
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, stat)
}
