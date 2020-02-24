package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func addCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
}

func middleWareHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// t := time.Now()
		// Add base headers
		addCORS(c)
		// Run next function
		c.Next()
		// // Log request
		// log.Infof("%v %v %v %s", c.Request.RemoteAddr, c.Request.Method, c.Request.URL, time.Since(t))
	}
}

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
	router.Use(middleWareHandler(), gin.Recovery())

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
