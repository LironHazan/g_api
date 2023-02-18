package main

import (
	"fmt"
	routes "g_api/libs/api/go-api/cmd"
	"g_api/libs/common-db/db"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func startServer(dbCtx db.DbContext) {
	app := fiber.New()
	routes.ResolveRoutes(app, dbCtx)
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
		dbCtx := db.DbContext{
			DB: _db,
		}
		startServer(dbCtx)
	}

}
