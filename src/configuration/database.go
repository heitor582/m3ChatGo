package configuration

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/heitor582/m3ChatGo/src/models"
)

var DBConn *gorm.DB

func InitDatabase() {
	var err error
	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable"
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	DBConn.AutoMigrate(&models.UserModel{}, &models.ChatRoomModel{}, &models.MessageModel{})
	fmt.Println("Database Migrate")
}