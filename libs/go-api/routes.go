package go_api

import (
	"g_api/libs/go-api/pkg/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// This function is for demo the nx executors
func Hello(name string) string {
	result := "Hello " + name
	return result
}

func ResolveRoutes(app *fiber.App, db *gorm.DB) {
	user.ResolveRoutes(app, db)
}
