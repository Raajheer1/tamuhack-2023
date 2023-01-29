package controllers

import (
	"fmt"
	"github.com/Raajheer1/tamuhack-2023/api/m/v2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserBookings(c *gin.Context) {
	fmt.Println("Getting user bookings...")
	user := c.MustGet("x-user").(*models.User)

	bookings, err := models.GetUserBookings(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}
