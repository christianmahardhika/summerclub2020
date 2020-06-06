// controllers/books.go

package controllers

import (
	"fmt"

	"github.com/christianmahardhika/summerclub2020/models"
	"github.com/gin-gonic/gin"
)

type CreateTweetInput struct {
	Username string `json:"username" binding:"required"`
	Post     string `json:"post" binding:"required"`
}

type UpdateTweetInput struct {
	Username string `json:"username"`
	Post     string `json:"post"`
}

// GET /tweets
// Find all tweets
func FindTweets(c *gin.Context) {
	var tweets []models.Twitter
	models.DB.Find(&tweets)

	if len(tweets) == 0 {
		c.JSON(404, gin.H{"success": false, "result": nil, "errorMessage": "Record not found!"})
	} else {
		c.JSON(200, gin.H{"success": true, "result": tweets, "errorMessage": nil})
	}
}

// GET /tweet/:id
// Find a twwet
func FindTweet(c *gin.Context) {
	// Get model if exist
	var tweet models.Twitter
	if err := models.DB.Where("id = ?", c.Query("id")).First(&tweet).Error; err != nil {
		c.JSON(404, gin.H{"success": false, "result": nil, "errorMessage": "Tweet Not Found"})
		return
	}

	c.JSON(200, gin.H{"success": true, "result": tweet, "errorMessage": nil})
}

// POST /tweets
// Create new tweet
func CreateTweet(c *gin.Context) {
	// Validate input
	var input CreateTweetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{"success": false, "result": nil, "errorMessage": err.Error()})
		return
	}

	// Create tweet
	tweet := models.Twitter{Username: input.Username, Post: input.Post}
	models.DB.Create(&tweet)

	c.JSON(201, gin.H{"success": true, "result": tweet, "errorMessage": nil})
}

// PATCH /tweet/:id
// Update a tweet
func UpdateTweet(c *gin.Context) {
	// Get model if exist
	var tweet models.Twitter
	if err := models.DB.Where("id = ?", c.Query("id")).First(&tweet).Error; err != nil {
		c.JSON(404, gin.H{"success": false, "result": nil, "errorMessage": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTweetInput
	fmt.Println(input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{"success": false, "result": nil, "errorMessage": err.Error()})
		return
	}

	models.DB.Model(&tweet).Updates(input)

	c.JSON(201, gin.H{"success": true, "result": tweet, "errorMessage": nil})
}

// DELETE /tweet/:id
// Delete a tweet
func DeleteTweet(c *gin.Context) {
	// Get model if exist
	var tweet models.Twitter
	if err := models.DB.Where("id = ?", c.Query("id")).First(&tweet).Error; err != nil {
		c.JSON(404, gin.H{"success": false, "result": nil, "errorMessage": "Record not found!"})
		return
	}

	models.DB.Delete(&tweet)

	c.JSON(201, gin.H{"success": true, "result": c.Query("id"), "errorMessage": nil})
}
