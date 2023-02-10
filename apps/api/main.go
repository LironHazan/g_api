package main

import (
	"fmt"
)

func main() {
	result := <-AsyncHashUserEmail(5)
	fmt.Println("This will be printed before the hashing will end!")
	result2 := <-AsyncHashUserEmail(4)
	fmt.Println("Probably this one as well")

	matchHash(result)
	matchHash(result2)
}
