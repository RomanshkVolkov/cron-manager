package handler

import (
	"net/http"

	"github.com/RomanshkVolkov/test-api/internal/core/service"
	"github.com/gin-gonic/gin"
)

func SyncElevaZapier(c *gin.Context) {
	response := service.SyncElevaZapier()

	c.IndentedJSON(http.StatusOK, response)
}

func SyncBoatyHealthChecker(c *gin.Context) {
	response := service.SyncBoatyHealthChecker()

	c.IndentedJSON(http.StatusOK, response)
}
