package controllers

import (
	"fmt"
	"net/http"

	"github.com/Raajheer1/tamuhack-2023/api/m/v2/models"
	utils "github.com/Raajheer1/tamuhack-2023/api/m/v2/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	fmt.Println("Login attempt")
	session := sessions.Default(c)

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	fmt.Println(input)

	user, err := models.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	session.Set("user-id", user.ID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(200, gin.H{"login": "successful"})
}

type SignupInput struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var input SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"e": err.Error()})
		return
	}

	u := models.User{}
	u.Email = input.Email
	u.Name = input.Name
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(400, gin.H{"e": err.Error()})
		return
	}
	u.Password = hashedPassword

	_, err = u.CreateUser()
	if err != nil {
		fmt.Println("Failed to create user")
		fmt.Println(err)
		c.JSON(400, gin.H{"e": err.Error()})
		return
	}

	c.JSON(200, gin.H{"signup": "successful"})
}
