package utils

import (
	"blog/models"
	"gorm.io/gorm"
)

func CreatePost(DB *gorm.DB, post *models.Post) error {
	err := DB.Create(post)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
func GetAllPosts(DB *gorm.DB) ([]models.Post, error) {
	var posts []models.Post
	err := DB.Find(&posts)
	if err.Error != nil {
		return nil, err.Error
	}
	return posts, nil
}
