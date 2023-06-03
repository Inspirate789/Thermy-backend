package server

import (
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server/middleware"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/authorization"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"regexp"
)

type Server struct {
	srv            *http.Server
	storageService storage.StorageManager
	authService    authorization.AuthManager
	logger         *log.Logger
}

func (s *Server) addRoutes(rg *gin.RouterGroup) {
	rg.POST("/units", s.postUnits)
	rg.PATCH("/units", s.patchUnits)
	rg.PUT("/units/all", s.getAllUnits)
	rg.PUT("/units/models", s.getUnitsByModels)
	rg.PUT("/units/properties", s.getUnitsByProperties)

	rg.GET("/models/all", s.getModels)
	rg.POST("/models", s.postModels)

	rg.GET("/elements/all", s.getModelElements)
	rg.POST("/elements", s.postElements)

	rg.PUT("/properties/all", s.getProperties)
	rg.PUT("/properties/unit", s.getPropertiesByUnit)
	rg.POST("/properties", s.postProperties)

	rg.GET("/layers/all", s.getAllLayers)
	rg.POST("/layers", s.postLayer)

	rg.POST("/users", s.postUser)
}

func parseRole(ctx *gin.Context) (string, error) {
	exp, err := regexp.Compile("/")
	if err != nil {
		return "", err
	}

	tokenIndex := len(exp.Split(os.Getenv("BACKEND_API_PREFIX"), -1))

	return exp.Split(ctx.FullPath(), -1)[tokenIndex], nil // "/.../role/..." --> "role"
}

func (s *Server) setupHandlers(router *gin.RouterGroup, authMgr authorization.AuthManager) {
	router.POST("/login", s.login)
	router.POST("/logout", s.logout)

	authRG := router.Group("", middleware.SessionCheck(s.authService))
	s.addRoutes(authRG)

	adminRg := authRG.Group("/admin", middleware.RoleCheck(authMgr, parseRole))
	adminRg.GET("/stat", s.getStat)
}

func NewServer(port int, authMgr authorization.AuthManager, storageMgr storage.StorageManager, logger *log.Logger) *Server {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()
	// router.SetTrustedProxies([]string{"192.168.52.38"}) // TODO?
	router.UseRawPath = true
	router.UnescapePathValues = false

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(gin.LoggerWithWriter(logger.Out))
	router.Use(middleware.ErrorResponseWriter(logger))
	router.Use(gin.RecoveryWithWriter(logger.Out))

	s := Server{ // TODO: Enabling SSL/TLS encryption
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		storageService: storageMgr,
		authService:    authMgr,
		logger:         logger,
	}

	apiRG := router.Group(os.Getenv("BACKEND_API_PREFIX"))
	s.setupHandlers(apiRG, authMgr)

	return &s
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
