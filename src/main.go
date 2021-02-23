package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	database "github.com/mathmed/challenge-bw-go/src/Infra/Database"
	routes "github.com/mathmed/challenge-bw-go/src/Routes"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	database.Config()
	routes.Setup(r)
	r.Run(":" + port)
	log.Printf("Listening on port %s", port)

}

