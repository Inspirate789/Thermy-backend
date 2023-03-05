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
	storageService storage.StorageManager
	authService    authorization.AuthManager
	log            logger.Logger
}

func (s *Server) addStudentRoutes(rg *gin.RouterGroup) {
	rg.POST("/units", s.postUnits)
	rg.PATCH("/units", s.patchUnits)
	rg.GET("/units/all", s.getAllUnits)
	rg.GET("/units/models", s.getUnitsByModels)
	rg.GET("/units/properties", s.getUnitsByProperties)

	rg.GET("/models/all", s.getModels)

	rg.GET("/elements/all", s.getModelElements)

	rg.GET("/properties/all", s.getProperties)
	rg.GET("/properties/unit", s.getPropertiesByUnit)
	rg.POST("/properties", s.postProperties)

	rg.GET("/layers/all", s.getAllLayers)
}

func (s *Server) addEducatorRoutes(rg *gin.RouterGroup) {
	rg.POST("/models", s.postModels)

	rg.POST("/elements", s.postElements)

	rg.POST("/layers", s.postLayer)
}

func (s *Server) addAdminRoutes(rg *gin.RouterGroup) {
	rg.POST("/user/create", s.postUser)
	rg.GET("/user/get/password", s.getUserPassword)
}

func parseRole(ctx *gin.Context) (string, error) {
	exp, err := regexp.Compile("/")
	if err != nil {
		return "", err
	}

	return exp.Split(ctx.FullPath(), 3)[1], nil // "/role/..." --> "role"
}

func (s *Server) setupHandlers(router *gin.Engine) {
	router.UseRawPath = true
	router.UnescapePathValues = false

	router.GET("/login", s.login)
	router.POST("/logout", s.logout)

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

func NewServer(port int, authMgr authorization.AuthManager, storageMgr storage.StorageManager, log logger.Logger) *Server {
	router := gin.Default()
	// router.SetTrustedProxies([]string{"192.168.52.38"}) // TODO?

	s := Server{ // TODO: Enabling SSL/TLS encryption
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		storageService: storageMgr,
		authService:    authMgr,
		log:            log,
	}
	s.setupHandlers(router)

	return &s
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
