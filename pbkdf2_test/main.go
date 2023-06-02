package main

import (
	"crypto/aes"
	"crypto/sha512"
	"fmt"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	fmt.Println("Hello world!")
	passwd := "foobar"
	salt := "baz"
	const derivedKeyLength = 2 * aes.BlockSize
	//derivedKey := pbkdf2.Key([]byte(passwd), salt, 100000, derivedKeyLength, sha512.New)

	start := time.Now()

	for i := 0; i < 100; i++ {
		fmt.Printf("on interation %v \n", i)
		pbkdf2.Key([]byte(passwd), []byte(salt), 100000, derivedKeyLength, sha512.New)
	}

	elapsed := time.Since(start)
	fmt.Printf("PBKDF2 took %s", elapsed)
}
