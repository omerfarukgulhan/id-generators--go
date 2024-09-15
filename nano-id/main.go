package main

import (
	"crypto/rand"
	"fmt"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Println(generateNanoID(21))
}

func generateNanoID(size int) string {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error generating random bytes:", err)
		return ""
	}

	id := make([]byte, size)
	for i := 0; i < size; i++ {
		id[i] = alphabet[int(b[i])%len(alphabet)]
	}

	return string(id)
}
