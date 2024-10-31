package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Msg  string
	Code int
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Health Check!",
		"code":    http.StatusOK,
	})
}
