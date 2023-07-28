// middleware/authorize.go

package middleware

import (
	"assign3/migrate"
	"assign3/models"
	"assign3/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	//"gorm.io/gorm"
	"net/http"
	"os"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the access token from the request (e.g., from the Authorization header)
		accessToken := c.GetHeader("Authorization")

		// If the access token is not present in the header, check if it's provided in a cookie
		if accessToken == "" {
			var err error
			accessToken, err = c.Cookie("accessToken")
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token not provided"})
				c.Abort()
				return
			}
		}

		// Validate the access token using the ValidateToken function from utils
		token, err := utils.ValidateToken(accessToken, os.Getenv("ACCESS_TOKEN_PUBLIC_KEY"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			c.Abort()
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

		// Set the user object in the context so that it can be accessed in subsequent middleware and request handlers
		c.Set("user", user)

		// Continue to the next middleware or request handler
		c.Next()
	}
}
