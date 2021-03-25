package costumers

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Get("/costumers/:document_id?", CostumersView.GetAndListView)
	app.Post("/costumers/", CostumersView.CreateView)
}
