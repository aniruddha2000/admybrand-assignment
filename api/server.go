package api

import (
	"fmt"
	"log"
	"os"

	"github.com/aniruddha2000/admybrand-assignment/api/controllers"
	"github.com/aniruddha2000/admybrand-assignment/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

// Initialize the database load the database and run the server
func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error getting env, not comming through %v", err)
	} else {
		fmt.Println("we are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	seed.Load(server.DB)
	server.Run(":8080")
}
