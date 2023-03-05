package server

import (
	"errors"
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
	username := ctx.Query("username")
	role := ctx.Query("role")

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = s.storageService.AddUser(conn, username, role)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": "ok"})
}

func (s *Server) getUserPassword(ctx *gin.Context) {
	token, err := strconv.ParseUint(ctx.Query("token"), 10, 64)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse token from URL"))
		return
	}
	username := ctx.Query("username")

	conn, err := s.authService.GetSessionConn(token)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	password, err := s.storageService.GetUserPassword(conn, username)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"password": password, "error": "ok"})
}
