// routes/authRoutes.go

package routes

import (
	"assign3/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/api/auth")

	authGroup.POST("/register", controllers.Register)
	authGroup.POST("/login", controllers.Login)
	authGroup.GET("/refresh", controllers.RefreshAccessToken)
	authGroup.POST("/logout", controllers.Logout)
}
