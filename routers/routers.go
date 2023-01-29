package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/controllers"
)

func Routing(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/category", controllers.CreateCategory)
	app.Get("/category", controllers.GetAllCategories)
	app.Get("/category/:id", controllers.GetCategoryByID)
	app.Put("/category/:id", controllers.UpdateCategory)
	app.Delete("/category/:id", controllers.DeleteCategory)

}
