package app

import (
	"GetfitWithPhysio-backend/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConfigDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:8889)/getfitwith_physio?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	// Handle If Error
	helper.HandleError(err)

	return db
}
