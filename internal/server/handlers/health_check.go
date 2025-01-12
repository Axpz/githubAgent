package handlers

import (
	"githubagent/internal/server/register"

	"github.com/gin-gonic/gin"
)

func init() {
	register.Register("GET /healthz", func(r *gin.Engine) {
		r.GET("/healthz", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "health",
			})
		})
	})

	register.Register("GET /ready", func(r *gin.Engine) {
		r.GET("/ready", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ready",
			})
		})
	})
}
