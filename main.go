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
		products := models.FetchProducts()

		return c.Render("index", fiber.Map{
			"orders":   orders,
			"tables":   tables,
			"products": products,
		})
	})

	router.Get("/orders", func(c *fiber.Ctx) error {
		orders := models.FetchOrders()

		return c.Render("orders", fiber.Map{
			"orders": orders,
		})
	})

	router.Post("/orders", func(c *fiber.Ctx) error {
		orderName := c.FormValue("order-name")
		orderType := models.Type(c.FormValue("type"))
		orderStatus := models.Status(c.FormValue("status"))
		orderItems := c.FormValue("items")
		orderCost := c.FormValue("cost")
		custName := c.FormValue("cust-name")
		custNumber := c.FormValue("cust-number")
		destination := c.FormValue("destination")
		datetime := c.FormValue("datetime")

		order := models.Order{
			Name:        orderName,
			Type:        orderType,
			Status:      orderStatus,
			Items:       orderItems,
			Cost:        orderCost,
			CustName:    custName,
			CustNumber:  custNumber,
			Destination: destination,
			DateTime:    datetime,
		}

		fmt.Println(order)
		return c.Render("partials/order_card", fiber.Map{
			"Name":        orderName,
			"Type":        orderType,
			"Status":      orderStatus,
			"Items":       orderItems,
			"Cost":        orderCost,
			"CustName":    custName,
			"CustNumber":  custNumber,
			"Destination": destination,
			"DateTime":    datetime,
		})
	})

	router.Post("/products", func(c *fiber.Ctx) error {
		productName := c.FormValue("product-name")
		productDesc := c.FormValue("product-desc")
		productPrice := c.FormValue("product-price")
		productImg := c.FormValue("product-img")

		product := models.Product{
			Name:        productName,
			Description: productDesc,
			Price:       productPrice,
			Image:       productImg,
		}

		fmt.Println(product)
		return c.Render("partials/product_card", fiber.Map{
			"Name":        productName,
			"Description": productDesc,
			"Price":       productPrice,
			"Image":       productImg,
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
