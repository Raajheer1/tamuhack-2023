package controllers

import (
	"fmt"
	"net/http"

	duffel "github.com/Raajheer1/tamuhack-2023/api/m/v2/duffel"
	models "github.com/Raajheer1/tamuhack-2023/api/m/v2/models"
	"github.com/gin-gonic/gin"
)

type BookFlightInput struct {
	OfferID string `json:"offer_id"`
}

func BookFlight(c *gin.Context) {
	fmt.Println("Booking flight...")
	var input BookFlightInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	offer, err := duffel.GetSingleOfferById(input.OfferID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user := c.MustGet("x-user").(*models.User)

	err = models.CreateFlight(user, offer.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, offer)
}
