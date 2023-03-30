package routes

import (
	todolist "todolist/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	app.Post("/api/Create", todolist.CreateArticle)
	app.Get("/api/Read", todolist.ReadArticles)
	app.Post("/api/Update/:name", todolist.UpdateArticle)
	app.Post("/api/Delete/:name", todolist.DeleteArticle)

}
