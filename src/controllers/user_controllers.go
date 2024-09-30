package controllers

import (
	"net/http"

	"github.com/fhorray/gingon/src/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var users = []models.User{
		{Name: "Francy", Email: "francy@email.com"},
}

	c.JSON(http.StatusOK, users)
}