package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, Base64 Encoding ")
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded : ", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Decoded : ", string(decoded))

	// to be URL safe , avoid '/' and '+'
	urlSafeEncoding := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URLSafeEncoding : ", urlSafeEncoding)

	// base 64 encoding
	// user cases - binary data transfer , data storage , embedding resources

}
