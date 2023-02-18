package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type DbRef struct {
	DB *gorm.DB
}

type DummyUser struct {
	Hash string `json:"hash"`
}

func (d DbRef) handler(c *fiber.Ctx) error {
	result := <-AsyncHashUserEmail(5)
	fmt.Println("This will be printed before the hashing will end!")
	result2 := <-AsyncHashUserEmail(4)
	fmt.Println("Probably this one as well")

	MatchHash(result)
	MatchHash(result2)

	user5 := DummyUser{}
	user5.Hash = fmt.Sprintf("%v", result)
	err := d.createUser(user5)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create user"})
	}

	user4 := DummyUser{}
	user5.Hash = fmt.Sprintf("%v", result)
	err = d.createUser(user4)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create user"})
	}

	return c.SendString("UP we go 👋!")
}

func (d DbRef) createUser(user DummyUser) error {
	if result := d.DB.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func ResolveRoutes(app *fiber.App, db *gorm.DB) {
	d := &DbRef{
		DB: db,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		err := d.handler(c)
		if err != nil {
			return err
		}
		return nil
	})
}
