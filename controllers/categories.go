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
		return c.SendStatus(400)
	}

	err = repositories.CreateCategory(database.DbConnection, *category)
	if err != nil {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"message": "Success insert category",
	})
}

func GetAllCategories(c *fiber.Ctx) error {

	categories, err := repositories.GetAllCategories(database.DbConnection)

	//jsonCategories, err := json.Marshal(categories)

	if err != nil {
		return c.SendStatus(400)
	} else {
		return c.JSON(categories)
	}
}

func GetCategoryByID(c *fiber.Ctx) error {
	idCategory, _ := strconv.Atoi(c.Params("id"))

	category, err := repositories.GetCategoryByID(database.DbConnection, idCategory)

	if err != nil {
		return c.SendStatus(400)
	} else {
		return c.JSON(category)
	}
}

func UpdateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	idCategory, _ := strconv.Atoi(c.Params("id"))

	err := c.BodyParser(&category)
	if err != nil {
		return c.SendStatus(400)
	}

	*category, err = repositories.UpdateCategory(database.DbConnection, idCategory, *category)
	if err != nil {
		return c.SendStatus(400)
	}

	return c.JSON(category)
}

func DeleteCategory(c *fiber.Ctx) error {
	idCategory, _ := strconv.Atoi(c.Params("id"))

	err := repositories.DeleteCategory(database.DbConnection, idCategory)

	if err != nil {
		return c.SendStatus(400)
	} else {
		return c.JSON(fiber.Map{
			"message": "1 category deleted",
		})
	}
}
