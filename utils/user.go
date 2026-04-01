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
func MyData(DB *gorm.DB, id uint) ([]models.User, error) {
	var user []models.User
	err := DB.Preload("Posts", nil).Preload("Tags", nil).Select("Name,Email").Where("id = ?", id).Find(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return user, nil
}
func ViewAllUsers(DB *gorm.DB) ([]models.User, error) {
	var users []models.User
	result := DB.Preload("Posts", nil).Preload("Tags", nil).Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("Error while viewing all users %v", result.Error)
	}
	fmt.Println("ID\t\tName\t\t\tEmail\t\t\tPosts\t\t\tTags")
	for _, u := range users {
		var tag interface{} = "N/A"
		var post interface{} = "N/A"
		if len(u.Tags) > 0 {
			tag = u.Tags[0].Name
		}
		if len(u.Posts) > 0 {
			post = u.Posts[0].Title
		}
		fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v\n", u.ID, u.Name, u.Email, post, tag)
	}
	return users, nil
}
