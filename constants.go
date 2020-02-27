package main

import (
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

type Constants struct{
    Port string
	StringPort string
    PostsPath string
}

var constants Constants

func loadConstants(){
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file loaded")
	}

	// configure server port
	port := os.Getenv("PORT")
	stringPort := fmt.Sprintf(":%v", port)
	postsPath := os.Getenv("POSTS_STORAGE_PATH")

	constants.Port = port
	constants.StringPort = stringPort
    constants.PostsPath = postsPath
	log.Println("LOADED CONSTANTS")
}

func (constantsP *Constants) getConstants(){
	constantsP.Port = constants.Port
	constantsP.StringPort = constants.StringPort
    constantsP.PostsPath = constants.PostsPath
}