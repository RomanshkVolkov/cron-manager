package http

import (
	"github.com/RomanshkVolkov/test-api/internal/adapters/handler"
	// "github.com/RomanshkVolkov/test-api/internal/adapters/middleware"
	"github.com/gin-gonic/gin"
)

func JobRoutes(r *gin.Engine) {
	// protect := middleware.Protected
	job := r.Group("/jobs")
	{
		job.GET("/eleva-zapier", handler.SyncElevaZapier)
	}
}
