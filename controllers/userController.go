// controllers/userController.go

package controllers

import (
	"assign3/migrate"
	"assign3/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMyText(c *gin.Context) {
	// Get the user ID from the authentication middleware
	userID := c.MustGet("userID").(uint)

	// Check if the user exists in the database using GORM
	var user models.User
	if err := migrate.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Retrieve the user's questions from the database
	var questions []models.Question
	if err := migrate.DB.Where("user_id = ?", userID).Find(&questions).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "questions not found"})
		return
	}

	// Return the questions as a response
	c.JSON(http.StatusOK, gin.H{"questions": questions})
}

type CreateTextRequest struct {
	Text string `json:"text" binding:"required"`
}

type CreateTextResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

func CreateText(c *gin.Context) {
	// Get the user ID from the request URL parameter
	userID := c.Param("id")

	// Retrieve the user from the database
	var user models.User
	if err := migrate.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Parse the request body and bind it to the CreateTextRequest struct
	var req CreateTextRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Create a new question based on the request data
	question := models.Question{
		Text:   req.Text,
		UserID: user.ID,
	}

	// Save the new question in the database
	if err := migrate.DB.Create(&question); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
		return
	}

	// Return the response with the created question's ID and text
	c.JSON(http.StatusCreated, CreateTextResponse{
		ID:   question.ID,
		Text: question.Text,
	})
}

func GetAll(c *gin.Context) {
	// Retrieve all questions from the database
	var questions []models.Question
	if err := migrate.DB.Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve questions"})
		return
	}
	// Return the questions as a response
	c.JSON(http.StatusOK, gin.H{"questions": questions})
}
