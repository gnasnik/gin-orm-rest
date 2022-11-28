package api

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gnasnik/gin-orm-rest/core/dao"
	"github.com/gnasnik/gin-orm-rest/core/logger"
	"github.com/gnasnik/gin-orm-rest/errors"
	"github.com/gnasnik/gin-orm-rest/generated/model"
	"github.com/gnasnik/gin-orm-rest/utils"
	"github.com/mssola/user_agent"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	loginStatusFailure = iota
	loginStatusSuccess
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func jwtGinMiddleware(secretKey string) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "User",
		Key:         []byte(secretKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginParams login
			if err := c.ShouldBind(&loginParams); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginParams.Username
			password := loginParams.Password
			userAgent := c.Request.Header.Get("User-Agent")
			ua := user_agent.New(userAgent)
			os := ua.OS()
			explorer, _ := ua.Browser()
			clientIP := utils.GetClientIP(c.Request)
			location := utils.GetLocationByIP(clientIP)

			user, err := loginByPassword(c.Request.Context(), userID, password)
			if err != nil {
				logger.AddLoginLog(&model.LoginLog{
					Ipaddr:        clientIP,
					Browser:       explorer,
					Os:            os,
					Status:        loginStatusFailure,
					Msg:           err.Error(),
					LoginLocation: location,
				})
				return nil, err
			}

			logger.AddLoginLog(&model.LoginLog{
				LoginUsername: userID,
				LoginLocation: location,
				Ipaddr:        clientIP,
				Browser:       explorer,
				Os:            os,
				Status:        loginStatusSuccess,
				Msg:           "success",
			})
			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*model.User); ok && v.Username == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
				"success": false,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}

func loginByPassword(ctx context.Context, username, password string) (interface{}, error) {
	user, err := dao.GetUserByUsername(ctx, username)
	if err != nil {
		log.Errorf("get user by username: %v", err)
		return nil, errors.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(password)); err != nil {
		log.Errorf("can't compare hash %s ans password %s: %v", user.PassHash, password, err)
		return nil, errors.ErrInvalidPassword
	}

	return &model.User{UUID: user.UUID, Username: user.Username, Role: user.Role}, nil
}
