package server

import (
	// "net/http"
	"github.com/gin-gonic/gin"
	store "CF-RSS/pkg/store"
)
func CreateRoutes() {
	client := store.OpenConnectionWithMongoDB()
	reader := client.Database("CF-RSS").Collection("recent-actions-final")
	router := gin.Default()

	router.GET("/activity/recent-actions", func(c *gin.Context) {
		getBlogs(c, reader)
	})
	router.POST("/user/login", func(c *gin.Context) {
		getBlogs(c, reader)
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