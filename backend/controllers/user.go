package controllers

import (
	"backend/entities"
	"backend/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{"username": username})
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	username := req.Username
	password := req.Password
	usecase := usecase.UserUsecase{}
	if err := usecase.CreateUser(username, password); err != nil {
		log.Println("Create user error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}

func GetGamesByUser(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{"username": username})
	loginUser := c.GetString("username")
	if loginUser != username {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	usecase := usecase.UserUsecase{}
	games, err := usecase.GetGamesByUser(username)
	if err != nil {
		log.Println("Get games error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, games)
}

func GetToken(c *gin.Context) {
	username := c.GetString("username")
	usecase := usecase.UserUsecase{}
	token, err := usecase.GetToken(username)
	if err != nil {
		log.Println("Get token error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateToken(c *gin.Context) {
	username := c.GetString("username")
	usecase := usecase.UserUsecase{}
	err := usecase.UpdateToken(username)
	if err != nil {
		log.Println("Set token error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}

func GetPermissionByUserGame(c *gin.Context) {
	game := c.Param("game")
	user := c.Param("user")
	if game == "" || user == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	usecase := usecase.UserUsecase{}
	permission, err := usecase.GetPermission(user, game)
	if err != nil {
		log.Println("Get permission error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"permission": permission})
}

type SetPermissionRequest struct {
	Game       string              `json:"game_id"`
	User       string              `json:"user_id"`
	Permission entities.Permission `json:"permission"`
}

func SetPermission(c *gin.Context) {
	var req SetPermissionRequest
	if err := c.BindJSON(&req); err != nil || req.Game == "" || req.User == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	username := c.GetString("username")
	if username == req.User {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You cannot set permission for yourself"})
		return
	}
	usecase := usecase.UserUsecase{}
	err := usecase.SetPermission(req.User, req.Game, req.Permission)
	if err != nil {
		log.Println("Set permission error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}

func GenerateSharedURL(c *gin.Context) {
	game := c.Param("game")
	if game == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	username := c.GetString("username")
	usecase := usecase.UserUsecase{}
	// check permission
	if !usecase.CheckPermission(username, game, entities.ADMIN) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	sharedURL, err := usecase.GenerateSharedURL(username, game)
	if err != nil {
		log.Println("Generate shared URL error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"shared_url": sharedURL})
}
