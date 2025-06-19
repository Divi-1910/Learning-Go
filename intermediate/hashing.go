package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	// 	key components
	// deterministic
	// fast computation
	// pre-image resistance
	// collision resistance
	// sha-256
	// sha-512
	// salting

	//  Best Practices
	// use of standard libraries
	// algorithm updates

	password := "password1234"
	hash := sha256.Sum256([]byte(password))
	fmt.Println("SHA-256 Hash : ", hash)
	fmt.Println("Password : ", password)
	fmt.Printf("SHA-256 hash hex value :  %x \n", hash)

	newhash := sha512.Sum512([]byte(password))
	fmt.Printf("SHA-512 hash hex value :  %x \n", newhash)
	fmt.Println("SHA-512 : ", newhash)

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error in generating salt : ", err)
		return
	}
	fmt.Println("Salt : ", salt)

	saltStr := base64.StdEncoding.EncodeToString(salt)
	hashedpassword := hashPassword(password, salt)
	fmt.Println("Hashed Password : ", hashedpassword)
	fmt.Println("Salt : ", saltStr)

	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error in decoding salt : ", err)
		return
	}
	fmt.Println("Decoded : ", string(decodedSalt))

	loginHash := hashPassword(password, decodedSalt)

	if loginHash == hashedpassword {
		fmt.Println("Correct password , welcome")
	} else {
		fmt.Println("Incorrect password , go away")
	}

}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil

}
