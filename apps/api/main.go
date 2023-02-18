package main

import (
	"fmt"
	"g_api/libs/common-db/db"
	routes "g_api/libs/go-api"
	"g_api/libs/go-api/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func startServer(db *gorm.DB) {
	app := fiber.New()
	routes.ResolveRoutes(app, db)
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	config := db.Config{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Password: viper.GetString("DB_PASS"),
		User:     viper.GetString("DB_USER"),
		SSLMode:  viper.GetString("DB_SSLMODE"),
		DBName:   viper.GetString("DB_NAME"),
	}

	fmt.Println(config.Host)

	if _db, err := db.Connection(config); err == nil {
		err := models.MigrateDummyUsers(_db)
		if err != nil {
			return
		}
		startServer(_db)
	}

}
