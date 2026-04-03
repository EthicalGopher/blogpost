package server

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"gorm.io/gorm"
)

type ServerDB struct {
	DB *gorm.DB
}

// Server initializes the Fiber app and registers all routes.
func Server(db *gorm.DB) *fiber.App {
	s := &ServerDB{DB: db}
	app := fiber.New()
	api := app.Group("/rest/api/v1")
	api.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 30 * time.Second,
		LimitReached: func(c fiber.Ctx) error {

			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	}))
	api.Get("/health", func(c fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// Public routes
	api.Post("/signup", s.SignUp)
	api.Get("/user", s.AllUsers)
	api.Post("/signin", s.SignIn)
	api.Post("/refresh", RefreshToken)
	api.Get("/post", s.AllPosts)
	api.Get("/post/:postid", s.GetPost)
	api.Get("/post/:postid/comment", s.GetComments)
	api.Get("/category", s.AllCategories)
	api.Get("/tag", s.AllTags)

	// Protected routes
	protected := api.Group("", JWTMiddleware)
	protected.Post("/category", s.CreateCategory)
	protected.Post("/tag", s.CreateTag)
	user := protected.Group("/user/:id", Verify)
	user.Get("/", s.AboutMe)
	user.Get("/post", s.MyPost)
	user.Post("/post", s.CreatePost)
	user.Put("/post/:postid", s.UpdatePost)
	user.Delete("/post/:postid", s.DeletePost)
	user.Post("/post/:postid/comment", s.CreateComment)
	user.Delete("/post/:postid/comment/:commentid", s.DeleteComment)

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404)
	})
	return app
}
