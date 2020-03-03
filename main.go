package main

import (
	Api "./api"
	Config "./config"
	DB "./db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var MongoClient *mongo.Client

func main() {
	// LOAD ENV VARS INTO GLOBAL CONSTANTS HANDLER
	Config.Configure()

	// CONFIGURE MONGO CONNECTION
	DB.Connect()

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
	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/ping", Api.Ping)
		apiRouter.GET("/posts", Api.GetPosts)
		apiRouter.GET("/post", Api.GetPost)
		apiRouter.POST("/post", Api.SavePost)
	}

	log.Printf("Starting server on %v with stringPort %v\n", Config.Constants.Port, Config.Constants.StringPort)
	printName("BLOG-GO")

	router.Run(Config.Constants.StringPort)
}
