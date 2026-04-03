package server

import (
	"blog/models"
	"blog/utils"
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type AuthResponse struct {
	UserID       uint   `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUp handles user registration.
func (s *ServerDB) SignUp(c fiber.Ctx) error {
	var user UserRequest
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	u := models.User{Name: user.Name, Email: user.Email, Password: user.Password}
	if err := s.DB.Create(&u).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}
	return c.SendStatus(201)
}

// SignIn handles user authentication.
func (s *ServerDB) SignIn(c fiber.Ctx) error {
	var user UserRequest
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	h := sha256.Sum256([]byte(user.Password))
	hashedPassword := hex.EncodeToString(h[:])

	var u models.User
	if err := s.DB.Where("email = ? AND password = ?", user.Email, hashedPassword).First(&u).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	accessToken, err := utils.GenerateAccessToken(u.ID, u.Email)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	refreshToken, err := utils.GenerateRefreshToken(u.ID, u.Email)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(AuthResponse{
		UserID:       u.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// RefreshToken handles refreshing the access token.
func RefreshToken(c fiber.Ctx) error {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.Bind().Body(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	claims, err := utils.ValidateToken(body.RefreshToken)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	newAccessToken, err := utils.GenerateAccessToken(claims.UserID, claims.Email)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(fiber.Map{
		"access_token": newAccessToken,
	})
}

// AllUsers, AboutMe... (rest remains similar but use snake_case JSON if needed)
func (s *ServerDB) AllUsers(c fiber.Ctx) error {
	var user []models.User
	user, err := utils.ViewAllUsers(s.DB)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(user)
}

func (s *ServerDB) AboutMe(c fiber.Ctx) error {
	idVal := c.Params("id")
	id, _ := strconv.Atoi(idVal)
	user, err := utils.AboutMe(s.DB, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user",
		})
	}
	return c.JSON(user)
}
