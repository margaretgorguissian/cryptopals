package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	buffer_one := "1c0111001f010100061a024b53535009181c"
	buffer_two := "686974207468652062756c6c277320657965"
	
	bytes_one, _ := hex.DecodeString(buffer_one)
	bytes_two, _ := hex.DecodeString(buffer_two)
	
	xored_bytes := make([]byte, len(bytes_one))
	
	for i := 0; i < len(bytes_one); i++{
		xored_bytes[i] = bytes_one[i] ^ bytes_two[i]
	} 
	xored_string := hex.EncodeToString(xored_bytes)
	
	if (xored_string == "746865206b696420646f6e277420706c6179"){
		fmt.Println(xored_string)
	}
}
