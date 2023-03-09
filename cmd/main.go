package main

import (
	firebase "fyfirman-blog-service/cmd/repository"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	allowedOrigin := os.Getenv("ALLOWED_ORIGINS")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{allowedOrigin}

	r.Use(cors.New(config))

	return r
}

func main() {
	firebase.Initialize(os.Getenv("FIREBASE_DATABASE_URL"))

	r := setupRouter()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pongs")
	})

	confServerPort := os.Getenv("SERVER_PORT")
	if confServerPort == "" {
		log.Fatal("SERVER_PORT config is required")
	}
	r.Run(":" + confServerPort)
}
