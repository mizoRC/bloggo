package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// LOAD ENV VARS INTO GLOBAL CONSTANTS HANDLER
	loadConstants()

	// SET GLOBAL MAIN CONSTANTS FROM GLOBAL HANDLER
	var constants Constants
	constants.getConstants()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	//CORS
	router.Use(cors.Default())

	//router.Static("/asset-manifest.json", "./blog-go-client/build/asset-manifest.json")
	//router.Static("/service-worker.js", "./blog-go-client/build/service-worker.js")
	//router.Static("/", "./blog-go-client/build/")
	router.Use(static.Serve("/", static.LocalFile("./blog-go-client/build/", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/ping", ping)
	}

	postsPath := getPostStoragePath()
	log.Printf("POSTS STORAGE PATH %v", postsPath)

	log.Printf("Starting server on %v with stringPort %v\n", constants.Port, constants.StringPort)
	printName("BLOG-GO")

	router.Run(constants.StringPort)
}
