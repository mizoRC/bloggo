package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Consts struct {
	Port       string
	StringPort string
}

var Constants Consts

func Configure() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file loaded")
	}

	// configure server port
	port := os.Getenv("PORT")
	stringPort := fmt.Sprintf(":%v", port)

	Constants.Port = port
	Constants.StringPort = stringPort
	log.Println("Loaded constants")
}
