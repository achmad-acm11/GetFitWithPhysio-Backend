package app

import (
	"GetfitWithPhysio-backend/helper"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConfigDB() *gorm.DB {
	// errEnv := godotenv.Load(".env")
	// if errEnv != nil {
	// 	panic("Failed to load env file. Make sure .env file is exists!")
	// }
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	configuration := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, db_name)
	db, err := gorm.Open(mysql.Open(configuration), &gorm.Config{})

	// Handle If Error
	helper.HandleError(err)

	return db
}
