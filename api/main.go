package main

import (
	"fmt"
	"log"
	"os"

	auth "github.com/Raajheer1/tamuhack-2023/api/m/v2/controllers/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	controllers "github.com/Raajheer1/tamuhack-2023/api/m/v2/controllers"
	models "github.com/Raajheer1/tamuhack-2023/api/m/v2/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	models.ConnectDatabase()
	// initalize gin router
	r := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(auth.Auth)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"m": "Hello, Soap!",
		})
	})

	authRoutes := r.Group("/auth")
	authRoutes.POST("/login", controllers.Login)
	authRoutes.POST("/signup", controllers.Signup)

	r.POST("/search", controllers.Search)
	r.GET("/user", auth.NotGuest, controllers.GetUserInfo)
	r.POST("/book", auth.NotGuest, controllers.BookFlight)
	r.GET("/bookings", auth.NotGuest, controllers.GetUserBookings)

	err := r.Run()
	if err != nil {
		fmt.Println("Server run failed")
		return
	}
}
