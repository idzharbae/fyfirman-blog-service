package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"fyfirman-blog-service/valueobject"

	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
)

func setupRouter(client *db.Client) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pongs")
	})

	r.GET("/blog/:slug", func(c *gin.Context) {
		slug := c.Params.ByName("slug")

		blogRef := client.NewRef("blogs/" + slug)
		var blog valueobject.Blog
		if err := blogRef.Get(context.Background(), &blog); err != nil {
			log.Fatalln("Error reading value:", err)
		}

		c.JSON(http.StatusOK, blog)
	})

	r.POST("/blog/:slug/read", func(c *gin.Context) {
		slug := c.Params.ByName("slug")

		blogRef := client.NewRef("blogs/" + slug)

		var oldBlog valueobject.Blog
		if err := blogRef.Get(context.Background(), &oldBlog); err != nil {
			log.Fatalln("Error reading value:", err)
		}
		var data map[string]interface{}

		if oldBlog.Slug == "" {
			data = map[string]interface{}{
				"slug":      slug,
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
				"readCount": 1,
			}
		} else {
			data = map[string]interface{}{
				"slug":      oldBlog.Slug,
				"createdAt": oldBlog.CreatedAt,
				"updatedAt": time.Now(),
				"readCount": oldBlog.ReadCount + 1,
			}
		}

		if err := blogRef.Set(context.TODO(), data); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, oldBlog)
	})

	return r
}

func main() {
	client := InitializeFirebase()

	r := setupRouter(client)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
