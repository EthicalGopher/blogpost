package utils

import (
	"blog/models"
	"fmt"

	"gorm.io/gorm"
)

// CreateUser creates a new user in the database.
func CreateUser(DB *gorm.DB, user *models.User) error {
	result := DB.Create(user)
	if result.Error != nil {
		return fmt.Errorf("Error while creating user %v", result.Error)
	}
	return nil
}
// MyData returns the user data for a given email and password.
func MyData(DB *gorm.DB, email string, password string) (models.User, error) {
	var user models.User
	err := DB.Preload("Posts", nil).Preload("Tags", nil).Where("Email = ?", email, "Password = ?", password).Find(&user)
	if err.Error != nil {
		return user, err.Error
	}
	if user.ID == 0 {
		return models.User{}, fmt.Errorf("no user found")
	}
	return user, nil
}
// AboutMe returns the user data for a given user ID.
func AboutMe(DB *gorm.DB, id uint) (models.User, error) {
	var user models.User
	err := DB.Preload("Posts", nil).Preload("Tags", nil).Where("ID = ?", id).Find(&user)
	if err.Error != nil {
		return user, err.Error
	}
	if user.ID == 0 {
		return models.User{}, fmt.Errorf("no user found")
	}
	return user, nil
}

// ViewAllUsers returns all users in the database with their posts and tags preloaded.
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
