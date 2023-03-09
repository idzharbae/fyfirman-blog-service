package blogController

import (
	"context"
	firebase "fyfirman-blog-service/cmd/repository"
	"fyfirman-blog-service/valueobject"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FindBlog(c *gin.Context) {
	slug := c.Params.ByName("slug")

	blogRef := firebase.DB.NewRef("blogs/" + slug)
	var blog valueobject.Blog
	if err := blogRef.Get(context.Background(), &blog); err != nil {
		log.Fatalln("Error reading value:", err)
	}

	c.JSON(http.StatusOK, blog)
}

func ReadBlog(c *gin.Context) {
	slug := c.Params.ByName("slug")

	blogRef := firebase.DB.NewRef("blogs/" + slug)

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
}
