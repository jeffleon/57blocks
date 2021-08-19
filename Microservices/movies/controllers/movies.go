package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffleon/57block-movies/database"
	"github.com/jeffleon/57block-movies/models"
	"strconv"
)

func GetMovies(c *fiber.Ctx) error {
	var moviesByUser []models.Movies
	database.DB.Where(&models.Movies{UserID: 1}).Find(&moviesByUser)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "peliculas asociadas al usuario", "data": moviesByUser})
}

func GetMovie(c *fiber.Ctx) error {
	var movie models.Movies
	id := c.Params("id")
	database.DB.First(&movie, id)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "pelicula especifica", "data": movie})
}

func NewMovie(c *fiber.Ctx) error {
	var movie models.Movies
	err := c.BodyParser(&movie)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Bad request", "data": err})
	}
	err = database.DB.Model(&models.Movies{}).Create(&movie).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Bad request", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Se ha creado exitosamente la pelicula", "data": movie})
}

func UpdateMovie(c *fiber.Ctx) error {
	var editMovie models.Movies
	var movie models.Movies
	id := c.Params("id")
	if err := c.BodyParser(&editMovie); err != nil {
		return c.Status(503).Send([]byte("error"))
	}
	i, _ := strconv.ParseInt(id, 0, 64)
	err := database.DB.First(&movie, i).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "success", "message": "No se ha creado la pelicula", "data": err})
	}
	if movie.Name == "" {
		return c.Status(404).Send([]byte("El usuario no fue encontrado"))
	}

	movie.Genre = editMovie.Genre
	movie.Name = editMovie.Name
	movie.Description = editMovie.Description
	movie.Director = editMovie.Director
	movie.Release = editMovie.Release
	movie.UserID = editMovie.UserID

	database.DB.Save(&movie)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Se ha actualizado exitosamente la pelicula", "data": movie})
}

func DeleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie models.Movies

	i, _ := strconv.ParseInt(id, 0, 64)
	err := database.DB.First(&movie, i).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "success", "message": "no se ha creado la pelicula", "data": err})
	}
	if movie.Name == "" {
		return c.Status(404).Send([]byte("la pelicula no fue encontrado"))
	}
	database.DB.Delete(&movie)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "se ha eliminado exitosamente el usuario", "data": movie})
}
