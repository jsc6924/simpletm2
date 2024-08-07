package controllers

import (
	"backend/entities"
	"backend/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryTranslatesByGame(c *gin.Context) {
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

type IUResult struct {
	Result  string `json:"Result"`
	Message string `json:"Message"`
}

func InsertOrUpdateTranslate(c *gin.Context) {
	// get game, rawWord, translate from REST url
	game := c.Param("game")
	rawWord := c.Param("rawWord")
	translate := c.Param("translate")
	if game == "" || rawWord == "" || translate == "" {
		c.JSON(http.StatusBadRequest, IUResult{"False", "invalid request"})
		return
	}

	username := c.GetString("username")
	u := usecase.UserUsecase{}
	if !u.CheckPermission(username, game, entities.EDIT) {
		c.JSON(http.StatusUnauthorized, IUResult{"False", "Unauthorized"})
		return
	}
	apiGameUsecase := usecase.ApiGameUsecase{}
	err := apiGameUsecase.InsertOrUpdateTranslate(game, rawWord, translate)
	if err != nil {
		log.Println("insert or update error: ", err)
		c.JSON(http.StatusInternalServerError, IUResult{"False", "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, IUResult{"True", ""})
}

type IURequest struct {
	Game      string `json:"game"`
	RawWord   string `json:"rawWord"`
	Translate string `json:"translate"`
}

func InsertOrUpdateTranslateV2(c *gin.Context) {
	// get game, rawWord, translate from json as IURequest
	var iuRequest IURequest
	if err := c.BindJSON(&iuRequest); err != nil || iuRequest.Game == "" || iuRequest.RawWord == "" || iuRequest.Translate == "" {
		c.JSON(http.StatusBadRequest, IUResult{"False", "invalid request"})
		return
	}

	username := c.GetString("username")
	u := usecase.UserUsecase{}
	if !u.CheckPermission(username, iuRequest.Game, entities.EDIT) {
		c.JSON(http.StatusUnauthorized, IUResult{"False", "Unauthorized"})
		return
	}
	apiGameUsecase := usecase.ApiGameUsecase{}
	err := apiGameUsecase.InsertOrUpdateTranslate(iuRequest.Game, iuRequest.RawWord, iuRequest.Translate)
	if err != nil {
		log.Println("insert or update error: ", err)
		c.JSON(http.StatusInternalServerError, IUResult{"False", "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, IUResult{"True", ""})
}

func DeleteTranslate(c *gin.Context) {
	// get game, rawWord from REST url
	game := c.Param("game")
	rawWord := c.Param("rawWord")
	if game == "" || rawWord == "" {
		c.JSON(http.StatusBadRequest, IUResult{"False", "invalid request"})
		return
	}

	username := c.GetString("username")
	u := usecase.UserUsecase{}
	if !u.CheckPermission(username, game, entities.EDIT) {
		c.JSON(http.StatusUnauthorized, IUResult{"False", "Unauthorized"})
		return
	}
	apiGameUsecase := usecase.ApiGameUsecase{}
	err := apiGameUsecase.DeleteTranslate(game, rawWord)
	if err != nil {
		log.Println("delete error: ", err)
		c.JSON(http.StatusInternalServerError, IUResult{"False", "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, IUResult{"True", ""})
}

type DeleteRequest struct {
	Game    string `json:"game"`
	RawWord string `json:"rawWord"`
}

func DeleteTranslateV2(c *gin.Context) {
	// get game, rawWord from json as DeleteRequest
	var deleteRequest DeleteRequest
	if err := c.BindJSON(&deleteRequest); err != nil || deleteRequest.Game == "" || deleteRequest.RawWord == "" {
		c.JSON(http.StatusBadRequest, IUResult{"False", "invalid request"})
		return
	}

	username := c.GetString("username")
	u := usecase.UserUsecase{}
	if !u.CheckPermission(username, deleteRequest.Game, entities.EDIT) {
		c.JSON(http.StatusUnauthorized, IUResult{"False", "Unauthorized"})
		return
	}
	apiGameUsecase := usecase.ApiGameUsecase{}
	err := apiGameUsecase.DeleteTranslate(deleteRequest.Game, deleteRequest.RawWord)
	if err != nil {
		log.Println("delete error: ", err)
		c.JSON(http.StatusInternalServerError, IUResult{"False", "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, IUResult{"True", ""})
}
