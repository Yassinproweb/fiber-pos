package main

import (
	"fmt"

	"github.com/Yassinproweb/fiber-pos/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	router := fiber.New(fiber.Config{
		Views: engine,
	})

	router.Static("/static", "./static")

	router.Get("/", func(c *fiber.Ctx) error {
		orders := models.FetchOrders()
		tables := models.FetchTables()

		return c.Render("index", fiber.Map{
			"orders": orders,
			"tables": tables,
		})
	})

	router.Post("/orders", func(c *fiber.Ctx) error {
		orderName := c.FormValue("order-name")
		orderType := models.Type(c.FormValue("type"))
		orderStatus := models.Status(c.FormValue("status"))
		orderItems := c.FormValue("items")
		orderCost := c.FormValue("cost")

		order := models.Order{
			Name:   orderName,
			Type:   orderType,
			Status: orderStatus,
			Items:  orderItems,
			Cost:   orderCost,
		}

		fmt.Println(order)
		return c.Render("order_row", fiber.Map{
			"Name":   orderName,
			"Type":   orderType,
			"Status": orderStatus,
			"Items":  orderItems,
			"Cost":   orderCost,
		})
	})

	router.Post("/tables", func(c *fiber.Ctx) error {
		tableName := c.FormValue("table-name")
		tableState := models.State(c.FormValue("state"))
		tableCapacity := c.FormValue("capacity")

		table := models.Table{
			Name:     tableName,
			State:    tableState,
			Capacity: tableCapacity,
		}

		fmt.Println(table)
		return c.Render("table_row", fiber.Map{
			"Name":     tableName,
			"State":    tableState,
			"Capacity": tableCapacity,
		})
	})

	fmt.Println("Server is running on port 5174")
	router.Listen(":5174")
}
