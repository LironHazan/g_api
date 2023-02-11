package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx) error {
	result := <-AsyncHashUserEmail(5)
	fmt.Println("This will be printed before the hashing will end!")
	result2 := <-AsyncHashUserEmail(4)
	fmt.Println("Probably this one as well")

	MatchHash(result)
	MatchHash(result2)

	return c.SendString("UP we go 👋!")
}
