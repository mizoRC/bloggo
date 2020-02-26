package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file loaded")
	}

	// configure server port
	port := os.Getenv("PORT")
	stringPort := fmt.Sprintf(":%v", port)

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

	log.Printf("Starting server on %v with stringPort %v\n", port, stringPort)
	printName("BLOG-GO")

	router.Run(stringPort)
}
