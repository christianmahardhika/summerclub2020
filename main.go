package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/christianmahardhika/summerclub2020/controllers"
	"github.com/christianmahardhika/summerclub2020/models"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/api/twitter/list", controllers.FindTweets)
	r.GET("/api/twitter/detail", controllers.FindTweet)
	r.POST("/api/twitter/create", controllers.CreateTweet)
	r.PATCH("/api/twitter/update", controllers.UpdateTweet)
	r.DELETE("/api/twitter/delete", controllers.DeleteTweet)

	r.Run(":5000")
}
