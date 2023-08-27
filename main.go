package main

import (
	"fmt"
	"server/config"
	"server/models"
	"server/router"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect: ", err)
	}

	db.AutoMigrate(&models.Task{})
	config.DB = db

	server := router.Setup()
	server.Run()
}
