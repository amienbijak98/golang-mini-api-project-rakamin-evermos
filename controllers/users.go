package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/database"
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/models"
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/repositories"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST new User",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	err = repositories.Register(database.DbConnection, *user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST new User",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Success to POST new user",
		"errors":  err,
		"data":    1,
	})
}
