package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"fmty/controllers"
	"fmty/models"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("dist", false)))
	r.Use(cors.Default())

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/api/message/list", controllers.GetMessages)
	r.POST("/api/message/post", controllers.PostMessage)

	r.Run(":5000")
}
