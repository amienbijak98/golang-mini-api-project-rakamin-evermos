package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/database"
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/models"
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/repositories"
)

func CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)

	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST new category",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	err = repositories.CreateCategory(database.DbConnection, *category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to POST new category",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "Success to POST new category",
		"errors":  err,
		"data":    1,
	})
}

func GetAllCategories(c *fiber.Ctx) error {
	categories, err := repositories.GetAllCategories(database.DbConnection)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET category",
			"errors":  err.Error(),
			"data":    categories,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":    categories,
			"status":  true,
			"message": "Success to GET category",
			"errors":  err,
		})
	}
}

func GetCategoryByID(c *fiber.Ctx) error {
	idCategory, _ := strconv.Atoi(c.Params("id"))

	category, err := repositories.GetCategoryByID(database.DbConnection, idCategory)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to GET category",
			"errors":  err.Error(),
			"data":    nil,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":    category,
			"status":  true,
			"message": "Success to GET category",
			"errors":  err,
		})
	}
}

func UpdateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	idCategory, _ := strconv.Atoi(c.Params("id"))

	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE category",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	*category, err = repositories.UpdateCategory(database.DbConnection, idCategory, *category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to UPDATE category",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    *category,
		"status":  true,
		"message": "Success to UPDATE category",
		"errors":  err,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	idCategory, _ := strconv.Atoi(c.Params("id"))

	err := repositories.DeleteCategory(database.DbConnection, idCategory)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to DELETE category",
			"errors":  err.Error(),
			"data":    nil,
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  true,
			"message": "Success to DELETE category",
			"errors":  err,
			"data":    1,
		})
	}
}
