package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Static("/", "../service/templates")

	log.Fatal(app.Listen(":3000"))
}
