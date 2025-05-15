package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type User struct {
	Name  string
	Email string
}

func main() {
	engine := html.New("./views", ".html")

	router := fiber.New(fiber.Config{
		Views: engine,
	})

	router.Static("/static", "./static")

	router.Get("/", func(c *fiber.Ctx) error {
		users := GetUsers()
		return c.Render("index", fiber.Map{
			"users": users,
		})
	})

	router.Post("/users", func(c *fiber.Ctx) error {
		user := User{
			Name:  c.FormValue("name"),
			Email: c.FormValue("email"),
		}
		return c.Render("user_row", user)
	})

	router.Delete("/users/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		fmt.Println("Delete user called ", name)
		return c.SendStatus(fiber.StatusOK)
	})

	fmt.Println("Server is running on port 5174")
	router.Listen(":5174")
}

func GetUsers() []User {
	return []User{
		{Name: "Katungi Yassin", Email: "yassinkatungi67@gmail.com"},
		{Name: "Nnansereko Hajarah", Email: "nnanserekohajarah95@gmail.com"},
	}
}
