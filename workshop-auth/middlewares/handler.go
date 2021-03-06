package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NoMethodHandler
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "Method not found"})
	}
}

// NoRouteHandler
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound,
			gin.H{"message": "The processing function of the request route was not found"})
	}
}
