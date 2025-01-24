package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/RomanshkVolkov/test-api/internal/adapters/repository"
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// middeleware for cors and jwt

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now().UTC()
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		origin := c.GetHeader("Origin")
		fmt.Println("origin", origin)

		user, err := repository.ExtractDataByToken(token)

		if err == nil && user.ID != 0 {
			c.Set("user", user)
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, ngrok-skip-browser-warning")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatusJSON(http.StatusForbidden, domain.APIResponse[string]{
				Message: domain.Message{
					En: "Forbidden",
					Es: "Запрещено",
				},
			})
			return
		}

		c.Next()

		latency := time.Since(t)
		fmt.Println(latency)

		status := c.Writer.Status()
		fmt.Println(status)

	}
}
