package utils

import (
	"blog/models"
	"fmt"

	"gorm.io/gorm"
)

func CreateUser(DB *gorm.DB, user *models.User) error {
	result := DB.Create(user)
	if result.Error != nil {
		return fmt.Errorf("Error while creating user %v", result.Error)
	}
	return nil
}

func ViewAllUsers(DB *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("Error while viewing all users %v", result.Error)
	}
	fmt.Println("ID\t\tName\t\t\tEmail\t\t\tPosts\t\t\tTags")
	for _, u := range users {
		var comments interface{} = "N/A"
		var tag interface{} = "N/A"

		if len(u.Posts) > 0 {
			comments = u.Posts[0].Comments
		}
		if len(u.Tags) > 0 {
			tag = u.Tags[0].Name
		}

		fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v\n", u.ID, u.Name, u.Email, comments, tag)
	}
	return users, nil
}
