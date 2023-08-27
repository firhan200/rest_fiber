package main

import (
	"github.com/firhan200/rest_fiber/db"
	"github.com/firhan200/rest_fiber/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

type Book struct {
	Name string
}

func main() {
	//auto migrate
	db.Migrate()

	viewEngine := django.New("./pages", ".django")

	app := fiber.New(fiber.Config{
		Views: viewEngine,
	})

	handlers.HandlerProducts(app)

	app.Listen(":8080")
}
