package controllers

import (
	"backend/entities"
	"backend/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersByGame(c *gin.Context) {
	game := c.Param("game")
	if game == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no game"})
		return
	}
	username := c.GetString("username")
	u := usecase.UserUsecase{}
	if !u.CheckPermission(username, game, entities.READ) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	gu := usecase.GameUsecase{}
	users, err := gu.GetUsersByGame(game)
	if err != nil {
		log.Println("Get users error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, users)
}

type CreateGameRequest struct {
	GameId    string `json:"game_id"`
	GameTitle string `json:"game_title"`
}

func CreateGameAsLoginUser(c *gin.Context) {
	var req CreateGameRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	gameId := req.GameId
	gameTitle := req.GameTitle
	username := c.GetString("username")

	gu := usecase.GameUsecase{}
	err := gu.CreateGameAsUser(gameId, gameTitle, username)
	if err != nil {
		log.Println("Create game error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Already exists or internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}

type DeleteGameRequest struct {
	GameId string `json:"game_id"`
}

func DeleteGame(c *gin.Context) {
	var req DeleteGameRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	gameId := req.GameId
	username := c.GetString("username")
	// check permission
	u := usecase.UserUsecase{}
	if !u.CheckPermission(username, gameId, entities.ADMIN) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	gu := usecase.GameUsecase{}
	err := gu.DeleteGame(gameId)
	if err != nil {
		log.Println("Delete game error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such game"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}
