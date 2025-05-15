package main

import (
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

	router.Get("/", func(c *fiber.Ctx) error {
		users := GetUsers()
		return c.Render("index", fiber.Map{
			"users": users,
		})
	})
}

func GetUsers() []User {
	return []User{
		{Name: "Katungi Yassin", Email: "yassinkatungi67@gmail.com"},
		{Name: "Nnassaazi Salwa", Email: "nnassaazisalwa95@gmail.com"},
	}
}
