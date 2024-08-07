package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{"username": username})
}
