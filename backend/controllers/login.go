package controllers

import (
	"backend/config"
	"backend/entities"
	"backend/models"
	"backend/utils"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func Login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Fetch the user from the database
	user, err := models.Users(models.UserWhere.ID.EQ(creds.Username)).One(context.Background(), config.DB)
	log.Println("User = %+v", user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized "})
		return
	}

	// Hash the provided password
	hashedPassword := utils.Hash(creds.Password)

	if hashedPassword != user.Salt {
		log.Println("hashed password not match")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &entities.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	// Set the JWT token as a cookie
	c.SetCookie("token", tokenString, int(time.Until(expirationTime).Seconds()), "/", config.SetCookieDomain, false, false)
	c.JSON(http.StatusOK, "ok")
}
