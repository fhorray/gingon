package controllers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/fhorray/gingon/src/models"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts = []models.Post{
		{Title: "Post XYZ", Author: "Francy Santos"},
}

	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read body"})
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))


	println("BODY: ", string(body))
}