package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// Generate a random byte slice
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Random bytes:", hex.EncodeToString(randomBytes))

	// Calculate the SHA256 hash of a message
	message := "Hello, world!"
	hash := sha256.Sum256([]byte(message))
	fmt.Println("SHA256 hash of message:", hex.EncodeToString(hash[:]))
}
