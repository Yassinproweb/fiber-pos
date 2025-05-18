package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Yassinproweb/fiber-pos/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./pos.db")
	if err != nil {
		log.Fatal(err)
	}

	tableQueries := []string{
		`CREATE TABLE IF NOT EXISTS products (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT,
      price REAL,
      stock INTEGER,
      image TEXT
      );`,
		`CREATE TABLE IF NOT EXISTS tables (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      booked BOOLEAN,
      capacity INTEGER
      );`,
		`CREATE TABLE IF NOT EXISTS orders (
      id TEXT PRIMARY KEY,
      table_id TEXT, 
      status TEXT, 
      created_at TEXT, 
      total_price REAL
      );`,
		`CREATE TABLE IF NOT EXISTS bought_items (
      order_id TEXT, 
      product_name TEXT,
      quantity INTEGER,
      price REAL
    );`,
	}
	for _, q := range tableQueries {
		if _, err := db.Exec(q); err != nil {
			log.Fatal(err)
		}
	}
	return db
}

func main() {
	db := initDB()

	handlerOrder := handlers.NewOrderHandler(db)
	handlerProduct := handlers.NewProductHandler(db)
	handlerTable := handlers.NewTableHandler(db)

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")

	// Order routes
	app.Get("/", handlerOrder.ListOrders)
	app.Post("/update-status", handlerOrder.UpdateStatus)
	app.Get("/receipt", handlerOrder.PrintReceipt)

	// Product routes
	app.Post("/products", handlerProduct.AddProduct)

	// Table routes
	app.Get("/tables/available", handlerTable.AvailableTables)

	fmt.Println("Server running at http://localhost:5174")
	log.Fatal(app.Listen(":5174"))

}
