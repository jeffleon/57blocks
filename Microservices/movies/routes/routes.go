package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffleon/57block-movies/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/user/:id/movies", controllers.GetMovies)
	app.Get("/user/movie/:id", controllers.GetMovie)
	app.Post("/user/movie", controllers.NewMovie)
	app.Post("/user/movie/:id", controllers.UpdateMovie)
	app.Delete("/user/movie/:id", controllers.DeleteMovie)
}
