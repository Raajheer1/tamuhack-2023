package controllers

import (
	"fmt"
	"github.com/Raajheer1/tamuhack-2023/api/m/v2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	fmt.Println("Getting user info...")

	userId := c.MustGet("x-id").(uint)

	user, err := models.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, user)
}
