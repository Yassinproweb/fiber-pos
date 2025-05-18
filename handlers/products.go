package handlers

import (
	"database/sql"

	// "github.com/Yassinproweb/fiber-pos/models"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	DB *sql.DB
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{DB: db}
}

func (h *ProductHandler) AddProduct(c *fiber.Ctx) error {
	name := c.FormValue("name")
	price := c.FormValue("price")
	stock := c.FormValue("stock")
	image := c.FormValue("image")

	_, err := h.DB.Exec("INSERT INTO products (name, price, stock, image) VALUES (?, ?, ?, ?)", name, price, stock, image)
	if err != nil {
		return c.Status(500).SendString("Could not add product")
	}
	return c.Redirect("/")
}

