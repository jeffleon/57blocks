package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffleon/57block/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)
	app.Post("/user/:id", controllers.UpdateUser)
	app.Get("/user/:id", controllers.GetUser)
}
