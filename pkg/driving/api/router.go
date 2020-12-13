package api

import (
	"github.com/gin-gonic/gin"

	"AmazingTalker/pkg/technical/logger"
)

func NewRouter(l logger.Level) *Router {
	switch l {
	case logger.TraceLevel, logger.DebugLevel:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	return &Router{router}
}

type Router struct {
	*gin.Engine
}
