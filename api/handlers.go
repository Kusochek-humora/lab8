// handlers.go
package api

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/your-username/your-project-name/models"
)

var (
	items       []models.Item
	nextItemID  = 1
	itemsMutex  sync.Mutex
)

// GetItems returns the list of items
func GetItems(c *fiber.Ctx) error {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	return c.JSON(items)
}

// AddItem adds a new item
func AddItem(c *fiber.Ctx) error {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	var newItem models.Item
	if err := c.BodyParser(&newItem); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newItem.ID = nextItemID
	nextItemID++
	items = append(items, newItem)

	return c.Status(http.StatusCreated).JSON(newItem)
}

// GetItemByID returns an item by ID
func GetItemByID(c *fiber.Ctx) error {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	for _, item := range items {
		if item.ID == id {
			return c.JSON(item)
		}
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
}

// RemoveItemByID removes an item by ID
func RemoveItemByID(c *fiber.Ctx) error {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return c.JSON(fiber.Map{"message": "Item deleted"})
		}
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
}

// GetItemsByValue returns items by value
func GetItemsByValue(c *fiber.Ctx) error {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	value := c.Params("value")
	var filteredItems []models.Item

	for _, item := range items {
		if item.Value == value {
			filteredItems = append(filteredItems, item)
		}
	}

	return c.JSON(filteredItems)
}

// UpdateItem updates an item
func UpdateItem(c *fiber.Ctx) error {
	itemsMutex.Lock()
	defer itemsMutex.Unlock()

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	property := c.Params("property")
	value := c.Params("value")

	for i, item := range items {
		if item.ID == id {
			switch property {
			case "name":
				items[i].Name = value
			case "value":
				items[i].Value = value
			default:
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid property"})
			}

			return c.JSON(items[i])
		}
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
}
