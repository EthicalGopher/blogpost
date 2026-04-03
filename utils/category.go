package utils

import (
	"blog/models"

	"gorm.io/gorm"
)

// CreateCategory creates a new category in the database.
func CreateCategory(DB *gorm.DB, category *models.Category) error {
	err := DB.Create(&category)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// GetAllCategory returns all categories in the database.
func GetAllCategory(DB *gorm.DB) ([]models.Category, error) {
	var categories []models.Category
	err := DB.Find(&categories)
	if err.Error != nil {
		return nil, err.Error
	}
	return categories, nil
}
