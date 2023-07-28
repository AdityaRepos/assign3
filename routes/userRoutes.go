// routes/userRoutes.go

package routes

import (
	"assign3/controllers"
	"assign3/middleware"
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/api/user")

	// Apply the authorize middleware to all routes in the userGroup
	userGroup.Use(middleware.Authorize())

	userGroup.GET("/getMyText/:id", controllers.GetMyText)
	userGroup.POST("/createText/:id", controllers.CreateText)

}
