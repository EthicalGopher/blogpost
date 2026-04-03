package server

import (
	"blog/models"
	"blog/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func (s *ServerDB) AllPosts(c fiber.Ctx) error {
	posts, err := utils.GetAllPosts(s.DB)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(posts)
}

func (s *ServerDB) GetPost(c fiber.Ctx) error {
	idStr := c.Params("postid")
	id, _ := strconv.Atoi(idStr)
	var post models.Post
	if err := s.DB.Joins("User").First(&post, id).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Status(200).JSON(post)
}

func (s *ServerDB) MyPost(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)
	posts, err := utils.GetPostsByUserID(s.DB, uint(id))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(posts)
}

func (s *ServerDB) CreatePost(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Post")
	}
	post := models.Post{Title: body.Title, Content: body.Content, UserID: uint(id)}
	if err := utils.CreatePost(s.DB, &post); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("%v", err))
	}
	return c.Status(201).JSON(post)
}

func (s *ServerDB) UpdatePost(c fiber.Ctx) error {
	postidStr := c.Params("postid")
	postid, _ := strconv.Atoi(postidStr)
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}
	var post models.Post
	if err := s.DB.First(&post, postid).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	post.Title = body.Title
	post.Content = body.Content
	if err := s.DB.Save(&post).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(post)
}

func (s *ServerDB) DeletePost(c fiber.Ctx) error {
	idStr := c.Params("postid")
	id, _ := strconv.Atoi(idStr)
	post, err := utils.DeletePost(s.DB, uint(id))
	if err != nil {
		return c.Status(404).SendString(fmt.Sprintf("%v", err))
	}
	return c.Status(200).JSON(post)
}

func (s *ServerDB) GetComments(c fiber.Ctx) error {
	postidStr := c.Params("postid")
	postid, _ := strconv.Atoi(postidStr)
	comments, err := utils.GetAllComments(s.DB, uint(postid))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(comments)
}

func (s *ServerDB) DeleteComment(c fiber.Ctx) error {
	commentidStr := c.Params("commentid")
	commentid, _ := strconv.Atoi(commentidStr)
	if err := utils.DeleteComment(s.DB, uint(commentid)); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(200)
}

func (s *ServerDB) CreateComment(c fiber.Ctx) error {
	postIdStr := c.Params("postid")
	postid, err := strconv.Atoi(postIdStr)
	if err != nil {
		return err
	}
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	var body struct {
		Text string `json:"text"`
	}
	if err := c.Bind().Body(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	comment := models.Comment{Text: body.Text, PostID: uint(postid), UserID: uint(id)}
	if err := s.DB.Create(&comment).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(201).JSON(comment)
}
