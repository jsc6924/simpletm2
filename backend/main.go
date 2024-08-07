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
	router.POST("/signup", controllers.CreateUser) //todo
	// Auth group
	auth := router.Group("/")
	auth.Use(midware.AuthMiddleware()) //todo
	{
		auth.GET("/logout", controllers.Logout)
		auth.GET("/me", controllers.Me)
		auth.GET("/user/:username/games", controllers.GetGamesByUser)
		auth.GET("/user/:username/token", controllers.GetToken)
		auth.POST("/user/:username/token", controllers.UpdateToken)
		auth.POST("/gameop/create", controllers.CreateGameAsUser)
		auth.POST("/gameop/delete", controllers.DeleteGame)
		auth.GET("/game/:name/translations", controllers.QueryTranslatesByGame)
		auth.GET("/game/:name/users", controllers.GetUsersByGame)
		auth.GET("/permission/:user/:game", controllers.GetPermissionByUserGame)
		auth.POST("/permission", controllers.SetPermission)
		auth.GET("/generateSharedURL/:game", controllers.GenerateSharedURL)
	}

	api := router.Group("/")
	api.Use(midware.ApiMiddleware())
	{
		api.GET("/api/querybygame/:game", controllers.QueryTranslatesByGame)
		api.GET("/api/insert/:game/:rawWord/:translate", controllers.InsertOrUpdateTranslate)
		api.GET("/api/update/:game/:rawWord/:translate", controllers.InsertOrUpdateTranslate)
		api.GET("/api/delete/:game/:rawWord", controllers.DeleteTranslate)

		api.GET("/api2/querybygame/:game", controllers.QueryTranslatesByGame)
		api.POST("/api2/update", controllers.InsertOrUpdateTranslateV2)
		api.POST("/api2/delete", controllers.DeleteTranslateV2)
	}

	router.Run("0.0.0.0:9000")
}
