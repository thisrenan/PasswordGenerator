package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSequence(n int) string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("!@#$%&*()-=+0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fmt.Println(randomSequence(10))
}
