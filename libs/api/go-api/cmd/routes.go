package routes

import (
	user "g_api/libs/api/go-api/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func ResolveRoutes(app *fiber.App) {
	app.Get("/", user.Handler)
}
