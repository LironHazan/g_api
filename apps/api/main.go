package main

import (
	routes "g_api/libs/api/go-api/cmd"
	"github.com/gofiber/fiber/v2"
)

func startServer() {
	app := fiber.New()
	routes.ResolveRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}

func main() {
	startServer()
}
