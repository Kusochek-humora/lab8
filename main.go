// main.go
package main

import (
	"github.com/gofiber/fiber"
	"github.com/your-username/your-project-name/api"
	"github.com/your-username/your-project-name/models"
	"sync"
)

var (
	itemsMutex sync.Mutex
)

func main() {
	app := fiber.New()

	app.Static("/static", "./static")

	apiGroup := app.Group("/api")
	{
		apiGroup.Get("/items", api.GetItems)
		apiGroup.Post("/items", api.AddItem)
		apiGroup.Get("/items/:id", api.GetItemByID)
		apiGroup.Delete("/items/:id", api.RemoveItemByID)
		apiGroup.Get("/items/search/:value", api.GetItemsByValue)
		apiGroup.Put("/items/:id/:property/:value", api.UpdateItem)
	}

	// Start the server
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
