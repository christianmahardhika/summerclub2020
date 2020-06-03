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
	r.GET("/api/todo/list", controllers.FindTodos)
	r.GET("/api/todo/detail", controllers.FindTodo)
	r.POST("/api/todo/create", controllers.CreateTodo)
	r.PATCH("/api/todo/update", controllers.UpdateTodo)
	r.DELETE("/api/todo/delete", controllers.DeleteTodo)

	r.Run(":5000")
}
