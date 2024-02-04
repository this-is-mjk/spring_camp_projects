package server

import (
	// "net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	store "CF-RSS/pkg/store"
)
func CreateRoutes() {
	client := store.OpenConnectionWithMongoDB()
	reader := client.Database("CF-RSS").Collection("recent-actions-final")
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/activity/recent-actions", func(c *gin.Context) {
		getBlogs(c, reader)
	})
	router.POST("/user/login", func(c *gin.Context) {
		
	})

	router.POST("/user/signup", func(c *gin.Context) {
		getBlogs(c, reader)
	})

	// router.Use(authenticateReq)

	router.GET("/user/activity/recent-actions", func(c *gin.Context) {
		getBlogs(c, reader)
	})

	router.POST("/user/blogs/subscribe", func(c *gin.Context) {
		getBlogs(c, reader)
	})

	router.POST("/user/blogs/unsubscribe", func(c *gin.Context) {
		getBlogs(c, reader)
	})
	
    router.Run("localhost:8080")

}