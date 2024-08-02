package main

import (
	"backend/config"
	"backend/controllers"
	"backend/midware"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	fmt.Println("JWT Secret = ", config.JwtKey)
	fmt.Println("CorsAllowOrigin = ", config.CorsAllowOrigin)

	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.CorsAllowOrigin}, // Replace with the allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin!",
		})
	})
	router.POST("/login", controllers.Login)
	// Auth group
	auth := router.Group("/")
	auth.Use(midware.AuthMiddleware())
	{
		auth.GET("/api/me", controllers.Me)
	}

	router.Run("0.0.0.0:9000")
}
