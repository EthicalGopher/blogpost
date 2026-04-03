package server

import (
	"blog/models"
	"blog/utils"

	"github.com/gofiber/fiber/v3"
)

// AllTags returns all tags.
func (s *ServerDB) AllTags(c fiber.Ctx) error {
	tags, err := utils.GetAllTags(s.DB)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(tags)
}

// CreateTag creates a new tag.
func (s *ServerDB) CreateTag(c fiber.Ctx) error {
	var tag models.Tag
	if err := c.Bind().Body(&tag); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}
	if err := utils.CreateTag(s.DB, &tag); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(201).JSON(tag)
}
