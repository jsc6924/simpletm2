package controllers

import (
	"backend/config"
	"backend/entities"
	"backend/usecase"
	"backend/utils"
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

func Login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userUsecase := usecase.UserUsecase{}
	hashedPassword := utils.Hash(creds.Password)
	ok, err := userUsecase.CheckPassword(creds.Username, hashedPassword)
	if !ok || err != nil {
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Password is incorrect")
		}
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

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", config.SetCookieDomain, false, false)
	c.JSON(http.StatusOK, "ok")
}
