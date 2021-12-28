package middleware

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

func ClientSet(clientset *kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("clientset", clientset)
		c.Next()
	}
}
