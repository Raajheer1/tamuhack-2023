package auth

import (
	"net/http"

	models "github.com/Raajheer1/tamuhack-2023/api/m/v2/models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user-id")
	if userId == nil {
		c.Set("x-guest", true)
		c.Next()
		return
	}

	user, err := models.GetUser(userId.(uint))
	if err == nil {
		c.Set("x-guest", false)
		c.Set("x-id", userId.(uint))
		c.Set("x-user", user)
		c.Set("x-auth-type", "cookie")
		c.Next()
		return
	}

	// If we get here, they had a cookie with an invalid user
	// so delete it.
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
