package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisedevguy/golang-mini-api-project-rakamin-evermos/controllers"
)

func Routing(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello! This is golang api mini project created by Bijak Amien Muttaqie. Github repo: https://github.com/wisedevguy/golang-mini-api-project-rakamin-evermos")
	})

	//=== eCommerce App ===
	api := app.Group("/api/v1")

	//-- Reg and Login/Logout Auth
	// api.Post("/auth/register", controllers.Register)
	// api.Post("/auth/login", controllers.Login)
	// api.Post("/auth/logout", controllers.Logout)

	//-- GET PROVINCE AND REGENCY
	//api.Get(/)

	//-- CATEGORY CRUD
	api.Post("/category", controllers.CreateCategory)
	api.Get("/category", controllers.GetAllCategories)
	api.Get("/category/:id", controllers.GetCategoryByID)
	api.Put("/category/:id", controllers.UpdateCategory)
	api.Delete("/category/:id", controllers.DeleteCategory)

}
