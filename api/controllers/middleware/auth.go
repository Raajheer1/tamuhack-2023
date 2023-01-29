package auth

import (
	"net/http"

	models "github.com/Raajheer1/tamuhack-2023/api/m/v2/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	session := sessions.Default(c)
	userId := 10
	//userId := session.Get("user-id")
	//if userId == nil {
	//	c.Set("x-guest", true)
	//	c.Next()
	//	return
	//}

	//user, err := models.GetUser(userId.(uint))
	var temp models.User
	temp.Name = "Raaj Patel"
	temp.Email = "raaj@raajpatel.dev"
	temp.ID = 10
	c.Set("x-guest", false)
	c.Set("x-id", userId)
	c.Set("x-user", temp)
	c.Set("x-auth-type", "cookie")
	c.Next()
	return

	// If we get here, they had a cookie with an invalid user
	// so delete it.
	return
	session.Delete("user-id")
	c.Set("x-guest", true)
	c.Next()
}

func NotGuest(c *gin.Context) {
	if c.GetBool("x-guest") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}
