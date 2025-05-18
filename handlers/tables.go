package handlers

import (
	"database/sql"

	"github.com/Yassinproweb/fiber-pos/models"
	"github.com/gofiber/fiber/v2"
)

type TableHandler struct {
	DB *sql.DB
}

func NewTableHandler(db *sql.DB) *TableHandler {
	return &TableHandler{DB: db}
}

func (h *TableHandler) AvailableTables(c *fiber.Ctx) error {
	rows, err := h.DB.Query("SELECT id, booked, capacity FROM tables WHERE booked = 0")
	if err != nil {
		return c.Status(500).SendString("Failed to fetch tables")
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var t models.Table
		err := rows.Scan(&t.ID, &t.Booked, &t.Capacity)
		if err != nil {
			continue
		}
		tables = append(tables, t)
	}
	return c.JSON(tables)
}
