package github

import (
	"githubagent/internal/server/register"

	"github.com/gin-gonic/gin"
)

func init() {
	register.Register("POST /github/webhook", func(r *gin.Engine) {
		r.POST("/github/webhook", webhookHandler)
	})
}
