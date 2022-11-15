package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gnasnik/titan-operator/config"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("api")

func ServerAPI(cfg *config.Config) {
	gin.SetMode(cfg.Mode)
	r := gin.Default()
	r.Use(cors.Default())

	apiV0 := r.Group("/api/v0")
	authMiddleware, err := jwtGinMiddleware(cfg.SecretKey)
	if err != nil {
		log.Fatalf("jwt auth middleware: %v", err)
	}

	err = authMiddleware.MiddlewareInit()
	if err != nil {
		log.Fatalf("authMiddleware.MiddlewareInit: %v", err)
	}

	apiV0.POST("/login", authMiddleware.LoginHandler)
	apiV0.POST("/logout", authMiddleware.LogoutHandler)
	apiV0.GET("/refresh_token", authMiddleware.RefreshHandler)

	apiV0.Use(authMiddleware.MiddlewareFunc())

	// dashboard-api
	apiV0.GET("/schedules", GetSchedulersHandler)

	if err := r.Run(cfg.ApiListen); err != nil {
		log.Fatalf("starting server: %v\n", err)
	}
}
