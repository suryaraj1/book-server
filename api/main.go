package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	app := fiber.New()

	// for health monitoring of server
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON("hello I am alive...")
	})

	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(books)
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
