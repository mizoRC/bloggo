package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// constants var is loaded from main.go file

func getPostStoragePath() string{
	return constants.PostsPath
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func writeMdFile(){
	constants.getConstants()
}