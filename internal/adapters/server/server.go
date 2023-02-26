package server

import (
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server/middleware"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/authorization"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/Inspirate789/Thermy-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type Server struct {
	srv            *http.Server
	storageService *storage.StorageService
	authService    *authorization.AuthorizationService
	log            logger.Logger
}

func (s *Server) addStudentRoutes(rg *gin.RouterGroup) {
	rg.POST("/units/:layer/add/:token", s.postUnits)
	rg.PATCH("/units/:layer/update/:token", s.patchUnits)
	rg.GET("/units/:layer/all/:token", s.getAllUnits)
	rg.GET("/units/:layer/models/:token", s.getUnitsByModels)
	rg.GET("/units/:layer/properties/:token", s.getUnitsByProperties)

	rg.GET("/models/:layer/:token", s.getModels)

	rg.GET("/elements/:layer/:token", s.getModelElements)

	rg.GET("/properties/:token", s.getProperties)
	rg.GET("/properties/unit/:layer/:token", s.getPropertiesByUnit)
	rg.POST("/properties/:token", s.postProperties)

	rg.GET("/layers/all/:token", s.getAllLayers)
}

func (s *Server) addEducatorRoutes(rg *gin.RouterGroup) {
	rg.POST("/models/:token", s.postModels)

	rg.POST("/elements/:token", s.postElements)

	rg.POST("/layer/:layer/:token", s.postLayer)
}

func (s *Server) addAdminRoutes(rg *gin.RouterGroup) {
	rg.POST("/user/:username/create/:role/:token", s.postUser)
	rg.GET("/user/:username/password/:token", s.getUserPassword)
}

func parseRole(ctx *gin.Context) (string, error) {
	exp, err := regexp.Compile("/")
	if err != nil {
		return "", err
	}

	return exp.Split(ctx.FullPath(), 3)[1], nil // "/role/..." --> "role"
}

func (s *Server) setupHandlers(router *gin.Engine) {
	router.GET("/login", s.login)
	router.POST("/logout/:token", s.logout)

	router.Use(middleware.ErrorHandler(s.log))

	studentRG := router.Group("/student")
	studentRG.Use(middleware.RoleCheck(s.authService, parseRole))
	s.addStudentRoutes(studentRG)

	educatorRG := router.Group("/educator")
	educatorRG.Use(middleware.RoleCheck(s.authService, parseRole))
	s.addStudentRoutes(educatorRG)
	s.addEducatorRoutes(educatorRG)

	adminRG := router.Group("/admin")
	adminRG.Use(middleware.RoleCheck(s.authService, parseRole))
	s.addStudentRoutes(adminRG)
	s.addEducatorRoutes(adminRG)
	s.addAdminRoutes(adminRG)
}

func NewServer(port int, authSVC *authorization.AuthorizationService, storageSVC *storage.StorageService, log logger.Logger) Server {
	router := gin.Default()
	// router.SetTrustedProxies([]string{"192.168.52.38"}) // TODO?

	s := Server{ // TODO: Enabling SSL/TLS encryption
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		storageService: storageSVC,
		authService:    authSVC,
		log:            log,
	}
	s.setupHandlers(router)

	return s
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
