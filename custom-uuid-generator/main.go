package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	customID := customUUID()
	fmt.Println("Custom UUID:", customID)
}

func customUUID() string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 64)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error generating random bytes:", err)
		return ""
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x-%x",
		timestamp,
		randomBytes[0:4],
		randomBytes[4:6],
		randomBytes[6:8],
		randomBytes[8:10],
		randomBytes[10:16],
	)
}
