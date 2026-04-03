package utils

import (
	"blog/models"

	"gorm.io/gorm"
)

// CreateTag creates a new tag in the database.
func CreateTag(DB *gorm.DB, tag *models.Tag) error {
	err := DB.Create(&tag)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// GetAllTags returns all tags in the database.
func GetAllTags(DB *gorm.DB) ([]models.Tag, error) {
	var tags []models.Tag
	err := DB.Find(&tags)
	if err.Error != nil {
		return nil, err.Error
	}
	return tags, nil
}
