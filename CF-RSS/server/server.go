package server

import (
	// "net/http"
	"github.com/gin-gonic/gin"
	store "CF-RSS/store"
)
func CreateRoutes() {
	client := store.OpenConnectionWithMongoDB()
	reader := client.Database("CF-RSS").Collection("recent-actions-final")
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		getBlogs(c, reader)
	})
 
    router.Run("localhost:8080")

}