package midware

import (
	"backend/config"
	"backend/usecase"
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "empty token"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "empty token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func ApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			log.Println("No Authorization header provided")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Decode the Base64 encoded credentials
		apiToken, err := base64.StdEncoding.DecodeString(authHeader[len("Basic "):])
		if err != nil {
			log.Println("Failed to decode API token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Split into username and password
		parts := strings.SplitN(string(apiToken), ":", 2)
		if len(parts) != 2 {
			log.Println("Invalid API token format")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		username, token := parts[0], parts[1]

		userUsecase := usecase.UserUsecase{}
		ok, err := userUsecase.CheckToken(username, token)
		if !ok || err != nil {
			log.Println("Invalid API token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Set the username in the context
		c.Set("username", username)
		c.Next()
	}
}
