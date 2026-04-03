package utils

import (
	"blog/models"
	"gorm.io/gorm"
)

// CreatePost creates a new post in the database.
func CreatePost(DB *gorm.DB, post *models.Post) error {
	err := DB.Create(post)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// GetPostsByUserID returns all posts by the user with the given ID.
func GetPostsByUserID(DB *gorm.DB, userID uint) ([]models.Post, error) {
	var posts []models.Post
	err := DB.Where("user_id = ?", userID).Find(&posts)
	if err.Error != nil {
		return nil, err.Error
	}
	return posts, nil
}

// GetAllPosts returns all posts in the database.
func GetAllPosts(DB *gorm.DB) ([]models.Post, error) {
	var posts []models.Post
	err := DB.Find(&posts)
	if err.Error != nil {
		return nil, err.Error
	}
	return posts, nil
}

// DeletePost deletes a post from the database.
func DeletePost(DB *gorm.DB, id uint) ([]models.Post, error) {
	var post []models.Post
	err := DB.Where("id = ?", id).Delete(&post)
	if err.Error != nil {
		return nil, err.Error
	}
	return post, nil
}

// CreateComment creates a new comment on a post.
func CreateComment(DB *gorm.DB, postid uint, userid uint, comment models.Comment) error {
	err := DB.Create(&comment)
	if err.Error != nil {
		return err.Error
	}
	return nil

}

// GetAllComments returns all comments for a post.
func GetAllComments(DB *gorm.DB, postid uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := DB.Where("post_id = ?", postid).Find(&comments)
	if err.Error != nil {
		return nil, err.Error
	}
	return comments, nil
}

// DeleteComment deletes a comment from the database.
func DeleteComment(DB *gorm.DB, commentid uint) error {
	var comment []models.Comment
	err := DB.Where("id = ?", commentid).Delete(&comment)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
