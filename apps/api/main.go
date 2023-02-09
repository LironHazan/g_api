package main

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func hash(input string) string {
	// Add a random salt to the input string
	rand.Seed(time.Now().UnixNano())
	salt := strconv.Itoa(rand.Int())
	input = input + salt

	// Perform SHA-512 hash on the input
	hash := sha512.New()
	hash.Write([]byte(input))
	sha := hex.EncodeToString(hash.Sum(nil))

	// Add the salt to the end of the hash
	sha = sha + salt

	// Perform another round of SHA-512 hash
	hash = sha512.New()
	hash.Write([]byte(sha))
	sha = hex.EncodeToString(hash.Sum(nil))

	// Return the final hash string
	return sha
}

func hashUserEmail(id int) (string, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/" + strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return "", err
	}
	// This is ugly :) but go's performance are too good and I'm trying to make a blocking computation
	return hash(hash(hash(hash(hash(hash(hash(hash(hash(hash(user.Email)))))))))), nil
}

func asyncHashUserEmail(id int) chan interface{} {
	start := time.Now()
	ch := make(chan interface{}, 1)
	go func() {
		_start := time.Now()
		user, err := hashUserEmail(id)
		elapsed := time.Since(_start)
		fmt.Println(elapsed)
		if err != nil {
			ch <- errors.New("failed")
		} else {
			ch <- user
		}
	}()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	return ch
}

func matchHash(result interface{}) {
	switch v := result.(type) {
	case string:
		fmt.Println(v)
	case error:
		fmt.Println(v)
	}
}

func main() {
	result := <-asyncHashUserEmail(5)
	result2 := <-asyncHashUserEmail(4)
	result3 := <-asyncHashUserEmail(3)

	matchHash(result)
	matchHash(result2)
	matchHash(result3)
}
