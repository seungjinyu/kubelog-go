package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"greetings": "welcome",
	})
}
func RedirectToWelcome(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/welcome")
}

func V1welcome(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"auth": true,
	})
}
