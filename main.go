package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Yassinproweb/fiber-pos/db"
	"github.com/Yassinproweb/fiber-pos/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Helper function to convert DATETIME to DD-MM-YYYY HH:MM
	convertFromSQLiteDateTime := func(sqliteDateTime string) (string, error) {
		parsed, err := time.Parse("2006-01-02 15:04:05", sqliteDateTime)
		if err != nil {
			return "", fmt.Errorf("invalid SQLite date_time: %v", err)
		}
		return parsed.Format("02-01-2006 15:04"), nil
	}

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

	router.Get("/tables", func(c *fiber.Ctx) error {
		tables := models.FetchTables()

		return c.Render("tables", fiber.Map{
			"tables": tables,
		})
	})

	router.Post("/orders", func(c *fiber.Ctx) error {
		orderName := c.FormValue("order-name")
		orderType := models.Type(c.FormValue("type"))
		orderStatus := models.Status(c.FormValue("status"))
		custName := c.FormValue("cust-name")
		custNumber := c.FormValue("cust-number")
		destination := c.FormValue("destination")
		dateTime := c.FormValue("datetime")
		tableIDStr := c.FormValue("table-id")      // New field for DineIn orders
		orderCartJSON := c.FormValue("order-cart") // JSON string for OrderCart

		// Convert table-id to int (optional for non-DineIn)
		var tableID sql.NullInt64
		if tableIDStr != "" {
			id, err := strconv.Atoi(tableIDStr)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid table-id")
			}
			tableID = sql.NullInt64{Int64: int64(id), Valid: true}
		}

		// Convert datetime to SQLite DATETIME format
		parsedDateTime, err := time.Parse("02-01-2006 15:04", dateTime)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid datetime format")
		}
		sqliteDateTime := parsedDateTime.Format("2006-01-02 15:04")

		// Parse OrderCart JSON
		var orderCart []models.OrderItem
		if orderCartJSON != "" {
			if err := json.Unmarshal([]byte(orderCartJSON), &orderCart); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid order-cart JSON")
			}
		}

		// Create Order struct
		order := models.Order{
			Name:        orderName,
			Type:        orderType,
			Status:      orderStatus,
			CustName:    custName,
			CustNumber:  custNumber,
			Destination: destination,
			DateTime:    dateTime, // Keep original format for rendering
			OrderCart:   orderCart,
		}

		// Validate and calculate Items and Cost
		if err := order.UpdateItemsAndCost(); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		// Insert into orders table
		result, err := db.DB.Exec(`INSERT INTO orders (order_name, type, status, cost, cust_name, cust_number, destination, date_time, table_id)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			orderName, orderType, orderStatus, order.Cost, custName, custNumber, destination, sqliteDateTime, tableID)
		if err != nil {
			log.Printf("Failed to insert order: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create order")
		}

		// Get the inserted order ID
		orderID, err := result.LastInsertId()
		if err != nil {
			log.Printf("Failed to get order ID: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve order ID")
		}

		// Insert OrderCart into order_items
		for _, item := range orderCart {
			_, err = db.DB.Exec(`INSERT INTO order_items (order_id, pdt_name, quantity)
				VALUES (?, ?, ?)`,
				orderID, item.PdtName, item.Quantity)
			if err != nil {
				log.Printf("Failed to insert order item: %v", err)
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to insert order items")
			}
		}

		// Fetch updated order (to get items and status after triggers)
		var updatedOrder models.Order
		var dbDateTime string
		err = db.DB.QueryRow(`SELECT order_name, type, status, items, cost, cust_name, cust_number, destination, date_time
			FROM orders WHERE id = ?`, orderID).
			Scan(&updatedOrder.Name, &updatedOrder.Type, &updatedOrder.Status, &updatedOrder.Items, &updatedOrder.Cost,
				&updatedOrder.CustName, &updatedOrder.CustNumber, &updatedOrder.Destination, &dbDateTime)
		if err != nil {
			log.Printf("Failed to fetch updated order: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch updated order")
		}

		// Convert DATETIME to model format
		updatedOrder.DateTime, err = convertFromSQLiteDateTime(dbDateTime)
		if err != nil {
			log.Printf("Failed to convert datetime: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to convert datetime")
		}

		// Fetch OrderCart
		rows, err := db.DB.Query(`SELECT pdt_name, quantity, unit_price FROM order_items WHERE order_id = ?`, orderID)
		if err != nil {
			log.Printf("Failed to fetch order items: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch order items")
		}
		defer rows.Close()

		updatedOrder.OrderCart = []models.OrderItem{}
		for rows.Next() {
			var item models.OrderItem
			if err := rows.Scan(&item.PdtName, &item.Quantity, &item.UnitPrice); err != nil {
				log.Printf("Failed to scan order item: %v", err)
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to scan order items")
			}
			updatedOrder.OrderCart = append(updatedOrder.OrderCart, item)
		}

		fmt.Println(updatedOrder)
		return c.Render("partials/order_card", fiber.Map{
			"Name":        updatedOrder.Name,
			"Type":        updatedOrder.Type,
			"Status":      updatedOrder.Status,
			"Items":       updatedOrder.Items,
			"Cost":        updatedOrder.Cost,
			"CustName":    updatedOrder.CustName,
			"CustNumber":  updatedOrder.CustNumber,
			"Destination": updatedOrder.Destination,
			"DateTime":    updatedOrder.DateTime,
			"OrderCart":   updatedOrder.OrderCart,
		})
	})

	router.Post("/products", func(c *fiber.Ctx) error {
		name := c.FormValue("product-name")
		desc := c.FormValue("product-desc")
		priceStr := c.FormValue("product-price")
		img := c.FormValue("product-img")

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid product price")
		}

		_, err = db.DB.Exec(`INSERT INTO products (name, description, price, image) VALUES (?, ?, ?, ?)`,
			name, desc, price, img)
		if err != nil {
			log.Printf("Failed to insert product: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create product")
		}

		return c.Render("partials/product_card", fiber.Map{
			"Name":        name,
			"Description": desc,
			"Price":       price,
			"Image":       img,
		})
	})

	router.Post("/tables", func(c *fiber.Ctx) error {
		tableName := c.FormValue("table-name")
		tableState := models.State(c.FormValue("state"))
		capacityStr := c.FormValue("capacity")

		capacity, err := strconv.Atoi(capacityStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid capacity")
		}

		_, err = db.DB.Exec(`INSERT INTO tables (table_name, capacity, state) VALUES (?, ?, ?)`,
			tableName, capacity, tableState)
		if err != nil {
			log.Printf("Failed to insert table: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create table")
		}

		return c.Render("table_row", fiber.Map{
			"Name":     tableName,
			"State":    tableState,
			"Capacity": capacity,
		})
	})

	fmt.Println("Server is running on port 5174")
	router.Listen(":5174")
}
