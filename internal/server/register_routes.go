package server

import (
	"githubagent/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	handlers.HealthCheckRegister(router)

	return router
}
