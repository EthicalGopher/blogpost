package tests

import (
	"net/http/httptest"
	"strings"
	"testing"

	"blog/models"
	"blog/server"
	"blog/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Comment{}, &models.Category{}, &models.Post{}, &models.Tag{})
	return db
}

func TestHealthCheck(t *testing.T) {
	db := SetupTestDB()
	app := server.Server(db)

	req := httptest.NewRequest("GET", "/rest/api/v1/health", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}
}

func TestAllPosts(t *testing.T) {
	db := SetupTestDB()
	app := server.Server(db)

	// Add dummy post
	db.Create(&models.Post{Title: "Test Post", Content: "Content"})

	req := httptest.NewRequest("GET", "/rest/api/v1/post", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}
}

func TestGetPost(t *testing.T) {
	db := SetupTestDB()
	app := server.Server(db)

	post := models.Post{Title: "Single Post", Content: "Only one"}
	db.Create(&post)

	req := httptest.NewRequest("GET", "/rest/api/v1/post/1", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d, body: %v", resp.StatusCode, resp.Body)
	}
}

func TestUpdatePost(t *testing.T) {
	db := SetupTestDB()
	app := server.Server(db)

	user := models.User{Name: "Test User", Email: "test@example.com"}
	db.Create(&user)

	post := models.Post{Title: "Old Title", Content: "Old Content", UserID: user.ID}
	db.Create(&post)

	token, _ := utils.GenerateAccessToken(user.ID, user.Email)

	updateBody := `{"title": "New Title", "content": "New Content"}`
	req := httptest.NewRequest("PUT", "/rest/api/v1/user/1/post/1", strings.NewReader(updateBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}
}

func TestGetComments(t *testing.T) {
	db := SetupTestDB()
	app := server.Server(db)

	post := models.Post{Title: "Post with comments", Content: "..."}
	db.Create(&post)

	db.Create(&models.Comment{PostID: post.ID, Text: "First comment"})
	db.Create(&models.Comment{PostID: post.ID, Text: "Second comment"})

	req := httptest.NewRequest("GET", "/rest/api/v1/post/1/comment", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}
}

func TestDeleteComment(t *testing.T) {
	db := SetupTestDB()
	app := server.Server(db)

	user := models.User{Name: "Commenter", Email: "commenter@example.com"}
	db.Create(&user)

	post := models.Post{Title: "Post", Content: "..."}
	db.Create(&post)

	comment := models.Comment{PostID: post.ID, UserID: user.ID, Text: "Delete me"}
	db.Create(&comment)

	token, _ := utils.GenerateAccessToken(user.ID, user.Email)

	req := httptest.NewRequest("DELETE", "/rest/api/v1/user/1/post/1/comment/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}
}

func TestCategory(t *testing.T) {
	db := SetupTestDB()
	app := server.Server(db)

	db.Create(&models.Category{Title: "Tech"})

	req := httptest.NewRequest("GET", "/rest/api/v1/category", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}
}

func TestTag(t *testing.T) {
	db := SetupTestDB()
	app := server.Server(db)

	db.Create(&models.Tag{Name: "Go"})

	req := httptest.NewRequest("GET", "/rest/api/v1/tag", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}
}
