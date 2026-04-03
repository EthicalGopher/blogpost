package server

import (
	"blog/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func JWTMiddleware(c fiber.Ctx) error {
	token := c.Cookies("access_token")

	if token == "" {
		authHeader := c.Get("Authorization")
		if authHeader != "" {
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) == 2 && tokenParts[0] == "Bearer" {
				token = tokenParts[1]
			}
		}
	}

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
	}

	claims, err := utils.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	// Store user info in context for next handlers
	c.Locals("user_id", claims.UserID)
	c.Locals("email", claims.Email)

	return c.Next()
}

func Verify(c fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if userID != uint(id) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
