package controllers

import (
	"backend/entities"
	"backend/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryByGame(c *gin.Context) {
	// get name form REST url
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
	apiGameUsecase := usecase.ApiGameUsecase{}
	translate, err := apiGameUsecase.QueryByGame(game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, translate)
}
