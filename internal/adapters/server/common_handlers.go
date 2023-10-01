package server

import (
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type loginResponse struct {
	Token string `json:"token"`
}

// login godoc
//
//	@Summary		Log in to the server.
//	@Description	log in to the server
//	@Tags			Auth
//	@Param			request	body	entities.AuthRequest	true	"Authentication request"
//	@Produce		json
//	@Success		200	{object}	loginResponse
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/login [post]
func (s *Server) login(ctx *gin.Context) { // TODO
	var request entities.AuthRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		s.logger.Error(err)
		_ = ctx.AbortWithError(http.StatusBadRequest, errors.ErrCannotParseJSONWrap("AuthRequest"))
		return
	}

	token, err := s.authService.AddSession(s.storageService, &request, ctx)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, loginResponse{Token: strconv.FormatUint(token, 10)})
}

// logout godoc
//
//	@Summary		Log out from the server.
//	@Description	log out from the server
//	@Tags			Auth
//	@Param			token	header	string	true	"User authentication token"
//	@Success		200
//	@Failure		400	{object}	string
//	@Failure		500	{object}	string
//	@Router			/logout [delete]
func (s *Server) logout(ctx *gin.Context) { // TODO

	ctx.Status(http.StatusOK)
}
