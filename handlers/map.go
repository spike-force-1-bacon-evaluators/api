package handlers

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// Map is a handler for the map page
func Map() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "map.tmpl", gin.H{
			"title": "bacon",
		})
	}
}
