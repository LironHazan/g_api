package routes

import (
	"g_api/libs/api/go-api/pkg/user"
	"g_api/libs/common-db/db"
	"github.com/gofiber/fiber/v2"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func ResolveRoutes(app *fiber.App, dbCtx db.DbContext) {
	app.Get("/", func(c *fiber.Ctx) error {
		err := user.Handler(c, dbCtx)
		if err != nil {
			return nil
		}
		return nil
	})
}
