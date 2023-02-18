package user

import (
	"fmt"
	"g_api/libs/api/go-api/pkg/models"
	"g_api/libs/common-db/db"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, dbCtx db.DbContext) error {
	result := <-AsyncHashUserEmail(5)
	fmt.Println("This will be printed before the hashing will end!")
	result2 := <-AsyncHashUserEmail(4)
	fmt.Println("Probably this one as well")

	MatchHash(result)
	MatchHash(result2)

	user5 := models.DummyUser{}
	user5.Hash = result
	err := createUser(dbCtx, user5)

	user4 := models.DummyUser{}
	user5.Hash = result2
	err = createUser(dbCtx, user4)

	if err != nil {
		return err
	}
	fmt.Println(dbCtx)
	// push to db
	// create := dbCtx.DB.Create(interface{})
	return c.SendString("UP we go 👋!")
}

func createUser(dbCtx db.DbContext, user models.DummyUser) error {
	err := dbCtx.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
