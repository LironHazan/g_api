package main

import (
	"fmt"
	user "g_api/apps/api/user"
)

func main() {
	result := <-user.AsyncHashUserEmail(5)
	fmt.Println("This will be printed before the hashing will end!")
	result2 := <-user.AsyncHashUserEmail(4)
	fmt.Println("Probably this one as well")

	user.MatchHash(result)
	user.MatchHash(result2)
}
