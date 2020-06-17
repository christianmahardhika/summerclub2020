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

}

// GET /todo/:id
// Find a todo
func FindTodo(c *gin.Context) {
	// Get model if exist

}

// POST /todos
// Create new todo
func CreateTodo(c *gin.Context) {
	// Validate input

	// Create todo

}

// PATCH /todo/:id
// Update a todo
func UpdateTodo(c *gin.Context) {
	// Get model if exist

	// Validate input

}

// DELETE /todo/:id
// Delete a todo
func DeleteTodo(c *gin.Context) {
	// Get model if exist

}
