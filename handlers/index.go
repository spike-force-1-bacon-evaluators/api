package handlers

import (
	"log"
	"net/http"

	"github.com/spike-force-1-bacon-evaluators/api/backend"
	"gopkg.in/gin-gonic/gin.v1"
)

// Restaurant wrap restaurant represetation
// for each row in ranking table
type Restaurant struct {
	Position  int
	Name      string
	Variation string
	Signal    string
}

// Index is a handler for the index page
func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		restaurantList, err := backend.Get()
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			log.Print(err)
		}
		// Serializes the struct as JSON into the response body
		// and also sets the Content-Type as "application/json".
		// c.JSON(http.StatusOK, restaurantSample())
		c.JSON(http.StatusOK, restaurantList)
	}
}
