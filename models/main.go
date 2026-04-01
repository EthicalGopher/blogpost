package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Posts    []Post `gorm:"foreignKey:UserID;references:ID"`
	Name     string `gorm:"size:40"`
	Email    string
	Password string
	Tags     []Tag `gorm:"many2many:all_tags"`
}

type Post struct {
	gorm.Model
	Title    string `gorm:"size:255"`
	Content  string
	UserID   uint
	Comments []Comment
}

type Category struct {
	gorm.Model
	PostID uint
	Title  string `gorm:"size:255"`
}
type Comment struct {
	gorm.Model
	PostID uint
	UserID uint
	Text   string
}
type Tag struct {
	gorm.Model
	Name string
}

func (user *User) AfterCreate(tx *gorm.DB) error {
	log.Println("Created user ", user.Name, "at id : ", user.ID)
	return nil
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	if user.Name == "" {
		return fmt.Errorf("Error name is empty")
	}
	if user.Email == "" && strings.Contains(user.Email, "@") {
		return fmt.Errorf("Error email is empty or not a valid email")
	}
	return nil
}
func (post *Post) AfterCreate(tx *gorm.DB) error {
	log.Println("Created Post Title : ", post.Title, "at id", post.ID)
	return nil
}
func (post *Post) BeforeCreate(tx *gorm.DB) error {
	if post.Title == "" || post.Content == "" {
		return fmt.Errorf("invalid post")
	}
	return nil

}
func (user *User) BeforeSave(tx *gorm.DB) error {
	h := sha256.Sum256([]byte(user.Password))
	hashedPassword := hex.EncodeToString(h[:])
	user.Password = hashedPassword
	return nil
}
