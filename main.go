// main.go

package main

import (
	"assign3/migrate"
	"assign3/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	migrate.MigrateDB()

	// Set up Gin router
	r := gin.Default()

	// Apply CORS middleware to allow requests from the frontend
	//r.Use(corsMiddleware())

	// Register routes
	routes.RegisterAuthRoutes(r)
	routes.RegisterUserRoutes(r)

	// Start the server
	r.Run(":8080")
}

/*package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createUser(c *gin.Context, db *gorm.DB) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&post)
	c.JSON(http.StatusCreated, post)
}

func getUsers(c *gin.Context, db *gorm.DB) {
	var posts []Post
	db.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func getUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var post Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func updateUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var post Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&post)
	c.JSON(http.StatusOK, post)
}

func deleteUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var post Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	db.Delete(&post)
	c.JSON(http.StatusNoContent, nil)
}

func main() {

	dsn := "host=arjuna.db.elephantsql.com user=cxnswktq password=uu9_6rMDxAxxXbCbc6K5UBpr3cU09yZ1 dbname=cxnswktq port=5432 sslmode=prefer"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Migrate the User model
	db.AutoMigrate(&Post{})
	// ...
	router := gin.Default()

	// CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// Create a user
	router.POST("/posts", func(c *gin.Context) { createUser(c, db) })

	// Retrieve all users
	router.GET("/posts", func(c *gin.Context) { getUsers(c, db) })

	// Retrieve a user by ID
	router.GET("/posts/:id", func(c *gin.Context) { getUser(c, db) })

	// Update a user by ID
	router.PUT("/posts/:id", func(c *gin.Context) { updateUser(c, db) })

	// Delete a user by ID
	router.DELETE("/posts/:id", func(c *gin.Context) { deleteUser(c, db) })

	// ...
	router.Run(":8080")
}
*/
