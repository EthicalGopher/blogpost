package server

import (
	"blog/models"
	"blog/utils"

	"github.com/gofiber/fiber/v3"
)

// AllCategories returns all categories.
func (s *ServerDB) AllCategories(c fiber.Ctx) error {
	categories, err := utils.GetAllCategory(s.DB)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(categories)
}

// CreateCategory creates a new category.
func (s *ServerDB) CreateCategory(c fiber.Ctx) error {
	var category models.Category
	if err := c.Bind().Body(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}
	if err := utils.CreateCategory(s.DB, &category); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(201).JSON(category)
}
