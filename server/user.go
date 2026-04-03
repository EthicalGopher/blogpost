package server

import (
	"blog/models"
	"blog/utils"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

type AuthResponse struct {
	UserID       uint   `json:"id"`
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
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	u := models.User{Name: user.Name, Email: user.Email, Password: user.Password}
	err := s.DB.Create(&u)
	if err.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(201)
}

// AllUsers returns all registered users.
func (s *ServerDB) AllUsers(c fiber.Ctx) error {
	var user []models.User
	var err error
	user, err = utils.ViewAllUsers(s.DB)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(user)
}

// SignIn handles user authentication.
func (s *ServerDB) SignIn(c fiber.Ctx) error {
	var user UserRequest
	if err := c.Bind().Body(&user); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	h := sha256.Sum256([]byte(user.Password))
	hashedPassword := hex.EncodeToString(h[:])

	// Correctly query the user
	var u models.User
	if err := s.DB.Where("email = ? AND password = ?", user.Email, hashedPassword).First(&u).Error; err != nil {
		return c.SendStatus(401)
	}

	accessToken, err := utils.GenerateAccessToken(u.ID, u.Email)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	refreshToken, err := utils.GenerateRefreshToken(u.ID, u.Email)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
	})

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

// AboutMe returns the profile of the currently logged-in user.
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
