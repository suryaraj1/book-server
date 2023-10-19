package main

import (
	"strconv"

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

	app.Get("/books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil || id < 1 || id > len(books) {
			return c.Status(404).JSON(fiber.Map{"error": "Book does not exist!"})
		}
		id -= 1 // not good practice
		return c.JSON(books[id])
	})

	app.Post("/books", func(c *fiber.Ctx) error {
		var newBook Book
		if err := c.BodyParser(&newBook); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		newBook.ID = len(books) + 1
		books = append(books, newBook)
		return c.Status(201).JSON(newBook)

	})

	app.Delete("/books", func(c *fiber.Ctx) error {
		books = nil
		return c.Status(204).Send(nil)
	})

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
