package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jeffleon/57block/database"
	"github.com/jeffleon/57block/models"
	"github.com/jeffleon/57block/utils"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const SecretKey = "57Blocks"

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	if !utils.ValidPassword(data["password"]) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "your password need to have a uppercase lower case character number and special character",
		})
	}
	database.DB.Create(&user)
	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	id := c.Params("id")
	database.DB.First(&user, id)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "usuario", "data": user})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(200).JSON(fiber.Map{
		"status":  "ok",
		"message": token,
		"id":      user.Id,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	var editUser models.User
	var user models.User
	id := c.Params("id")
	if err := c.BodyParser(&editUser); err != nil {
		return c.Status(503).Send([]byte("error"))
	}
	i, _ := strconv.ParseInt(id, 0, 64)
	err := database.DB.First(&user, i).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "success", "message": "No se ha encontrado el usuario", "data": err})
	}
	if user.Name == "" {
		return c.Status(404).Send([]byte("El usuario no fue encontrado"))
	}

	user.Name = editUser.Name
	user.Password = editUser.Password
	user.Email = editUser.Email

	database.DB.Save(&user)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Se ha actualizado exitosamente el usuario", "data": user})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}
