package main

import (
	"github.com/gin-gonic/gin"

	"github.com/christianmahardhika/summerclub2020/controllers"
	"github.com/christianmahardhika/summerclub2020/models"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/todos", controllers.FindTodos)
	r.GET("/todo/:id", controllers.FindTodo)
	r.POST("/todos", controllers.CreateTodo)
	r.PATCH("/todo/:id", controllers.UpdateTodo)
	r.DELETE("/todo/:id", controllers.DeleteTodo)

	r.Run(":5000")
}
