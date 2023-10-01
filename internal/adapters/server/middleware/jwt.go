package middleware

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"time"
)

type LoginResponse struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

type UserCheckFunc func(entities.AuthRequest) (*entities.User, error)

const (
	identityKey = "id"
	roleKey     = "role"
)

func MakeGinJWTMiddleware(requiredRoles []string, getUser UserCheckFunc) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "jwt",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) { // for login (step 1)
			var request entities.AuthRequest
			if err := c.ShouldBind(&request); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			return getUser(request)
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims { // for login (step 2)
			if v, ok := data.(*entities.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
					roleKey:     v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} { // for auth (step 1)
			claims := jwt.ExtractClaims(c)
			return &entities.User{
				ID:   int(claims[identityKey].(float64)),
				Role: claims[roleKey].(string),
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool { // for auth (step 2)
			v, ok := data.(*entities.User)
			return ok && slices.Contains(requiredRoles, v.Role)
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, LoginResponse{
				Token:  token,
				Expire: expire.Format(time.RFC3339),
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, LoginResponse{
				Token:  token,
				Expire: expire.Format(time.RFC3339),
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.Status(code)
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
