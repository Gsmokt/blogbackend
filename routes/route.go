package routes

import (
	"github.com/Gsmokt/blogbackend/controller"
	"github.com/Gsmokt/blogbackend/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Use(middleware.IsAuthenticate)
	app.Post("/api/post", controller.CreatePost)
	// app.Use(middleware.IsAuthenticate)
}
