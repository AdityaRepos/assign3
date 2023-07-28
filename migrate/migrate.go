package migrate

import (
	"assign3/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func MigrateDB() {
	dsn := "host=arjuna.db.elephantsql.com user=cxnswktq password=uu9_6rMDxAxxXbCbc6K5UBpr3cU09yZ1 dbname=cxnswktq port=5432 sslmode=prefer"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Migrate the User model
	DB.AutoMigrate(&models.User{})
}

func Loadenv(envkey string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(envkey)
}
