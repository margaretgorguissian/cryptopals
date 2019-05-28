package main

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
)

func main() {
	hex_string := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println("Hex: ", hex_string)
	bytes, _ := hex.DecodeString(hex_string)
	base64_string := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println("Base64: ", base64_string)
}
