// controllers/authController.go

package controllers

import (
	"assign3/migrate"
	"assign3/models"
	"assign3/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	//"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

// Implement functions for user registration, login, access token refresh, and logout.

// User registration
func Register(c *gin.Context) {
	// Parse the JSON request body into the Register struct
	var registerData models.Register
	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(registerData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Create a new User object
	newUser := models.User{
		Name:     registerData.Name,
		Email:    registerData.Email,
		Password: hashedPassword,
		Role:     registerData.Role,
		Branch:   registerData.Branch,
	}

	// Save the new user to the database using GORM
	migrate.DB.Create(&newUser)

	// Respond with success message
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// User login
func Login(c *gin.Context) {
	// Parse the JSON request body into the Login struct
	var loginData models.Login
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Retrieve the user from the database using the provided email
	var user models.User
	if err := migrate.DB.First(&user, loginData.Email).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Compare the provided password with the stored hashed password
	if !utils.ComparePassword(user.Password, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate and return the access token using the CreateToken function from utils
	accessToken, err := utils.CreateToken(time.Minute*15, user, os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	// Generate and set the refresh token
	refreshToken, err := utils.CreateToken(time.Hour*24*7, user, os.Getenv("REFRESH_TOKEN_PRIVATE_KEY"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	// Set the refresh token as an HTTPOnly cookie in the response header
	c.SetCookie("refreshToken", refreshToken, int((time.Hour * 24 * 7).Seconds()), "/", "", false, true)

	// Return the access token in the response body
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}

// Access token refresh
func RefreshAccessToken(c *gin.Context) {
	// Get the refresh token from the request (e.g., from a cookie)
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token not provided"})
		return
	}

	// Validate the refresh token using the ValidateToken function from utils
	token, err := utils.ValidateToken(refreshToken, os.Getenv("REFRESH_TOKEN_PUBLIC_KEY"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Get the user ID from the token claims
	userID := token.Claims.(jwt.MapClaims)["sub"].(uint)

	var user models.User

	// Check if the user exists in the database using GORM
	if err := migrate.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Generate and return a new access token using the CreateToken function from utils
	// Generate and set the access token
	accessToken, err := utils.CreateToken(time.Minute*15, user, os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	// Return the access token in the response body
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}

// User logout
func Logout(c *gin.Context) {
	// Expire the access and refresh token by setting it to an empty string and setting the MaxAge to 0
	c.SetCookie("accessToken", "", -1, "/", "", false, true)
	c.SetCookie("refreshToken", "", -1, "/", "", false, true)

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
