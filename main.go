package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Type string
type Status string

const (
	Takeaway Type = "Takeaway"
	Delivery Type = "Delivery"
	DineIn   Type = "DineIn"

	Placed    Status = "Placed"
	Pending   Status = "Pending"
	Ready     Status = "Ready"
	Canceled  Status = "Canceled"
	Delivered Status = "Delivered"
	Taken     Status = "Taken"
	Served    Status = "Served"
)

type Order struct {
	Name   string
	Type   Type
	Status Status
	Items  string
	Cost   string
}

func main() {
	engine := html.New("./views", ".html")

	router := fiber.New(fiber.Config{
		Views: engine,
	})

	router.Static("/static", "./static")

	router.Get("/", func(c *fiber.Ctx) error {
		orders := FetchOrders()
		return c.Render("index", fiber.Map{
			"orders": orders,
		})
	})

	router.Post("/orders", func(c *fiber.Ctx) error {
		order := Order{
			Name:   c.FormValue("name"),
			Type:   Type(c.FormValue("type")),
			Status: Status(c.FormValue("status")),
			Items:  c.FormValue("items"),
			Cost:   c.FormValue("cost"),
		}
		fmt.Println(order)
		return c.Render("order", order)
	})

	fmt.Println("Server is running on port 5174")
	router.Listen(":5174")
}

func FetchOrders() []Order {
	return []Order{
		{"#ORD0011", DineIn, Placed, "2", "7.33"},
		{"#ORD0012", Takeaway, Taken, "1", "3.85"},
		{"#ORD0013", DineIn, Canceled, "7", "22.00"},
		{"#ORD0014", Delivery, Pending, "8", "102.79"},
		{"#ORD0015", DineIn, Served, "5", "33.5"},
		{"#ORD0016", Delivery, Delivered, "13", "123.84"},
		{"#ORD0017", Takeaway, Ready, "2", "13.75"},
		{"#ORD0018", DineIn, Pending, "3", "23.05"},
	}
}
