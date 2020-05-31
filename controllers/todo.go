// controllers/books.go

package controllers

import (
	"fmt"

	"github.com/christianmahardhika/summerclub2020/models"
	"github.com/gin-gonic/gin"
)

type CreateTodoInput struct {
	Title string `json:"title" binding:"required"`
	Note  string `json:"note" binding:"required"`
}

type UpdateTodoInput struct {
	Title string `json:"title"`
	Note  string `json:"note"`
}

// GET /todos
// Find all todos
func FindTodos(c *gin.Context) {
	var todos []models.Todo
	models.DB.Find(&todos)

	if len(todos) == 0 {
		c.JSON(404, gin.H{"error": "Todo list not found"})
	} else {
		c.JSON(200, gin.H{"data": todos})
	}
}

// GET /todo/:id
// Find a todo
func FindTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found!"})
		return
	}

	c.JSON(200, gin.H{"data": todo})
}

// POST /todos
// Create new todo
func CreateTodo(c *gin.Context) {
	// Validate input
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Create todo
	todo := models.Todo{Title: input.Title, Note: input.Note}
	models.DB.Create(&todo)

	c.JSON(201, gin.H{"data": todo})
}

// PATCH /todo/:id
// Update a todo
func UpdateTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTodoInput
	fmt.Println(input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&todo).Updates(input)

	c.JSON(201, gin.H{"data": todo})
}

// DELETE /todo/:id
// Delete a todo
func DeleteTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&todo)

	c.JSON(201, gin.H{"data": "Deleted"})
}
