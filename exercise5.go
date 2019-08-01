package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	bytes :=  []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	encrypted := make([]byte, 0)
	for b := range bytes{
		encrypted = append(encrypted, bytes[b] ^ key[b % len(key)])	

	}
	fmt.Print(hex.EncodeToString(encrypted))
}
