package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mathmed/challenge-bw-go/src/Routes"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	Routes.Setup(r)
	r.Run(":" + port)
	log.Printf("Listening on port %s", port)

}

