package utils

import (
	"blog/models"
	"gorm.io/gorm"
)

func CreatePost(DB *gorm.DB, post *models.Post) error {
	return DB.Create(post).Error
}

func GetPostsByUserID(DB *gorm.DB, userID uint) ([]models.Post, error) {
	var posts []models.Post
	err := DB.Joins("User").Where("posts.user_id = ?", userID).Find(&posts).Error
	return posts, err
}

func GetAllPosts(DB *gorm.DB) ([]models.Post, error) {
	var posts []models.Post
	err := DB.Joins("User").Find(&posts).Error
	return posts, err
}

func DeletePost(DB *gorm.DB, id uint) ([]models.Post, error) {
	var post []models.Post
	err := DB.Where("id = ?", id).Delete(&post).Error
	return post, err
}

func CreateComment(DB *gorm.DB, postid uint, userid uint, comment models.Comment) error {
	comment.PostID = postid
	comment.UserID = userid
	return DB.Create(&comment).Error
}

func GetAllComments(DB *gorm.DB, postid uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := DB.Preload("User").Where("post_id = ?", postid).Find(&comments).Error
	return comments, err
}

func DeleteComment(DB *gorm.DB, commentid uint) error {
	return DB.Where("id = ?", commentid).Delete(&models.Comment{}).Error
}
