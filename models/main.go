package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Base model to replace gorm.Model with snake_case JSON tags
type Base struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	Base
	Posts    []Post `gorm:"foreignKey:UserID;references:ID" json:"posts"`
	Name     string `gorm:"size:40" json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Tags     []Tag  `gorm:"many2many:all_tags" json:"tags"`
}

type Post struct {
	Base
	Title    string    `gorm:"size:255" json:"title"`
	Content  string    `json:"content"`
	UserID   uint      `json:"user_id"`
	User     User      `gorm:"foreignKey:UserID" json:"author"`
	Comments []Comment `json:"comments"`
}

type Category struct {
	Base
	PostID uint   `json:"post_id"`
	Title  string `gorm:"size:255" json:"title"`
}

type Comment struct {
	Base
	PostID uint   `json:"post_id"`
	UserID uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
	Text   string `json:"text"`
}

type Tag struct {
	Base
	Name string `json:"name"`
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
	if user.Password != "" && len(user.Password) != 64 {
		h := sha256.Sum256([]byte(user.Password))
		hashedPassword := hex.EncodeToString(h[:])
		user.Password = hashedPassword
	}
	return nil
}
