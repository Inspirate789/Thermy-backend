package server

import (
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server/middleware"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

type Server struct {
	srv            *http.Server
	storageService StorageService
	logger         *log.Logger
}

func (s *Server) addStudentRoutes(rg *gin.RouterGroup) {
	rg.POST("/units", s.postUnits)
	rg.PATCH("/units", s.patchUnits)
	rg.GET("/units", s.getUnits)
	rg.PUT("/units/models", s.getUnitsByModels)
	rg.PUT("/units/properties", s.getUnitsByProperties)
	rg.DELETE("/units", s.deleteUnits)

	rg.GET("/models", s.getModels)

	rg.GET("/elements", s.getModelElements)

	rg.GET("/properties", s.getProperties)
	rg.PUT("/properties/unit", s.getPropertiesByUnit)
	rg.POST("/properties", s.postProperties)

	rg.GET("/layers", s.getLayers)
}

func (s *Server) addEducatorRoutes(rg *gin.RouterGroup) {
	rg.POST("/models", s.postModels)

	rg.POST("/elements", s.postElements)

	rg.POST("/layers", s.postLayer)
}

func (s *Server) addAdminRoutes(rg *gin.RouterGroup) {
	rg.POST("/users", s.postUser)
	rg.GET("/stat", s.getStat)
}

func (s *Server) setupJWTMiddleware(requiredRoles []string, getUser middleware.UserCheckFunc) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := middleware.MakeGinJWTMiddleware(requiredRoles, getUser)
	if err != nil {
		return nil, err
	}
	err = authMiddleware.MiddlewareInit()
	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}

func (s *Server) setupHandlers(router *gin.RouterGroup) error {
	authMiddleware, err := s.setupJWTMiddleware([]string{
		entities.StudentRole,
		entities.EducatorRole,
		entities.AdminRole,
	}, s.storageService.GetUser)
	if err != nil {
		return err
	}
	educatorMiddleware, err := s.setupJWTMiddleware([]string{
		entities.EducatorRole,
		entities.AdminRole,
	}, s.storageService.GetUser)
	if err != nil {
		return err
	}
	adminMiddleware, err := s.setupJWTMiddleware([]string{
		entities.AdminRole,
	}, s.storageService.GetUser)
	if err != nil {
		return err
	}

	router.POST("/login", s.login(authMiddleware.LoginHandler))
	authRg := router.Group("", authMiddleware.MiddlewareFunc())
	authRg.GET("/refresh", s.refresh(authMiddleware.RefreshHandler))
	authRg.DELETE("/logout", s.logout(authMiddleware.LogoutHandler))

	studentRg := authRg.Group("")
	s.addStudentRoutes(studentRg)

	educatorRg := authRg.Group("", educatorMiddleware.MiddlewareFunc())
	s.addEducatorRoutes(educatorRg)

	adminRg := authRg.Group("", adminMiddleware.MiddlewareFunc())
	s.addAdminRoutes(adminRg)

	return nil
}

func NewServer(port int, storageMgr StorageService, logger *log.Logger) (*Server, error) {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()
	router.UseRawPath = true
	router.UnescapePathValues = false

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(gin.LoggerWithWriter(logger.Out))
	router.Use(middleware.ErrorResponseWriter(logger))
	router.Use(gin.RecoveryWithWriter(logger.Out))

	s := Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		storageService: storageMgr,
		logger:         logger,
	}

	apiRG := router.Group(os.Getenv("BACKEND_API_PREFIX"))

	return &s, s.setupHandlers(apiRG)
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
