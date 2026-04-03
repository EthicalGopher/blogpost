package server

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// AllPosts returns all posts in the database.
func (s *ServerDB) AllPosts(c fiber.Ctx) error {
	posts, err := utils.GetAllPosts(s.DB)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(posts)
}

// GetPost returns a single post by ID.
func (s *ServerDB) GetPost(c fiber.Ctx) error {
	idStr := c.Params("postid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	var post models.Post
	if err := s.DB.First(&post, id).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(200).JSON(post)
}

// MyPost returns all posts by the user with the given ID.
func (s *ServerDB) MyPost(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	user, err := utils.GetPostsByUserID(s.DB, uint(id))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(user)
}

// CreatePost creates a new post in the database.
func (s *ServerDB) CreatePost(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	var Post struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.Bind().Body(&Post); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Post")
	}
	post := models.Post{Title: Post.Title, Content: Post.Content, UserID: uint(id)}
	err = utils.CreatePost(s.DB, &post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("%v", err))
	}
	return c.Status(201).JSON(post)

}

// UpdatePost updates an existing post in the database.
func (s *ServerDB) UpdatePost(c fiber.Ctx) error {
	postidStr := c.Params("postid")
	postid, err := strconv.Atoi(postidStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Post ID")
	}

	var updateData struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.Bind().Body(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	var post models.Post
	if err := s.DB.First(&post, postid).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	post.Title = updateData.Title
	post.Content = updateData.Content

	if err := s.DB.Save(&post).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(post)
}

// DeletePost deletes a post from the database.
func (s *ServerDB) DeletePost(c fiber.Ctx) error {
	idStr := c.Params("postid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	post, err := utils.DeletePost(s.DB, uint(id))
	if err != nil {
		return c.Status(404).SendString(fmt.Sprintf("%v", err))
	}
	return c.Status(200).JSON(post)
}

// GetComments returns all comments for a post.
func (s *ServerDB) GetComments(c fiber.Ctx) error {
	postidStr := c.Params("postid")
	postid, err := strconv.Atoi(postidStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Post ID")
	}

	comments, err := utils.GetAllComments(s.DB, uint(postid))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(comments)
}

// DeleteComment deletes a comment from the database.
func (s *ServerDB) DeleteComment(c fiber.Ctx) error {
	commentidStr := c.Params("commentid")
	commentid, err := strconv.Atoi(commentidStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Comment ID")
	}

	if err := utils.DeleteComment(s.DB, uint(commentid)); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(200)
}

// CreateComment creates a new comment on a post.
func (s *ServerDB) CreateComment(c fiber.Ctx) error {
	postIdStr := c.Params("postid")
	postid, err := strconv.Atoi(postIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	comment := models.Comment{Text: c.FormValue("text")}
	err = utils.CreateComment(s.DB, uint(postid), uint(id), comment)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(201)
}
