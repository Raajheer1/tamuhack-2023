package controllers

import (
	"fmt"
	"net/http"

	duffel "github.com/Raajheer1/tamuhack-2023/api/m/v2/duffel"
	"github.com/gin-gonic/gin"
)

type SearchInput struct {
	DepartureDay      string `json:"departure_day" binding:"required"`
	ReturnDay         string `json:"return_day"`
	FlightDestination string `json:"flight_destination" binding:"required"`
	FlightOrigin      string `json:"flight_origin" binding:"required"`
}

func Search(c *gin.Context) {
	fmt.Println("Searching flights...")

	var input SearchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if input.ReturnDay == "" {
		flights, err := duffel.SearchOneWay(input.FlightOrigin, input.FlightDestination, input.DepartureDay)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, flights)
		return
	}

	if input.ReturnDay == "" {
		input.ReturnDay = input.DepartureDay
	}
	flights, err := duffel.Search(input.FlightOrigin, input.FlightDestination, input.DepartureDay, input.ReturnDay)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, flights)

}
