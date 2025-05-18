package handlers

import (
	"database/sql"
	"time"

	"github.com/Yassinproweb/fiber-pos/models"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	DB *sql.DB
}

func NewOrderHandler(db *sql.DB) *OrderHandler {
	return &OrderHandler{DB: db}
}

func (h *OrderHandler) ListOrders(c *fiber.Ctx) error {
	rows, err := h.DB.Query("SELECT id, table_id, status, created_at, total_price FROM orders")
	if err != nil {
		return c.Status(500).SendString("Database error")
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		err := rows.Scan(&o.ID, &o.TableID, &o.Status, &o.CreatedAt, &o.TotalPrice)
		if err != nil {
			continue
		}
		orders = append(orders, o)
	}
	return c.Render("index", fiber.Map{"Orders": orders})
}

func (h *OrderHandler) UpdateStatus(c *fiber.Ctx) error {
	orderID := c.FormValue("order_id")
	status := c.FormValue("status")
	_, err := h.DB.Exec("UPDATE orders SET status = ? WHERE id = ?", status, orderID)
	if err != nil {
		return c.Status(500).SendString("Failed to update status")
	}
	return c.Redirect("/")
}

func (h *OrderHandler) PrintReceipt(c *fiber.Ctx) error {
	orderID := c.Query("order_id")
	var order models.Order
	order.ID = orderID
	order.CreatedAt = time.Now().Format("02 Jan 2006 15:04")

	rows, err := h.DB.Query("SELECT product_name, quantity, price FROM bought_items WHERE order_id = ?", orderID)
	if err != nil {
		return c.Status(500).SendString("DB error")
	}
	defer rows.Close()

	var total float64
	for rows.Next() {
		var item models.BoughtItem
		if err := rows.Scan(&item.ProductName, &item.Quantity, &item.Price); err != nil {
			continue
		}
		total += float64(item.Quantity) * item.Price
		order.Items = append(order.Items, item)
	}
	order.TotalPrice = total
	return c.Render("receipt", fiber.Map{"Order": order})
}
