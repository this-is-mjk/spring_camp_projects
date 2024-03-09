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
	userData := client.Database("CF-RSS").Collection("user")
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	// config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"content-type"}
	router.Use(cors.New(config))

	router.POST("/signup", func(c *gin.Context) {
		Signup(c, userData)
	})
	router.POST("/login", func(c *gin.Context) {
		loginUser(c, userData)
	})
	// middleware
	router.Use(authenticateReq)
	// secure
	router.GET("/all-recent-actions", func(c *gin.Context) {
		getBlogs(c, reader, userData)
	})
	router.POST("/subscribe-request", func(c *gin.Context) {
		SubscribeRequest(c, userData)
	})
    router.Run("localhost:8080")
}