package db

import (
	"blog/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connection() *gorm.DB {
	DB, err = gorm.Open(sqlite.Open("./repository/blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")
	return DB
}

func Migrate() {
	err = DB.AutoMigrate(&models.User{}, &models.Comment{}, &models.Category{}, &models.Post{}, &models.Tag{})
	if err != nil {
		log.Println("Error during migration", err)
	}
	log.Println("Created tables")
}
