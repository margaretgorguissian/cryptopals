package main

import (
	"encoding/hex"
	"fmt"
)
// Lazy version where I don't bother checking letter frequency, and
// instead just pick out the most sensical decrypted message

func main() {
	xored_string := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	xored_bytes, _ := hex.DecodeString(xored_string)
	decrypted := make([]byte, len(xored_bytes))
	for char := []byte("!")[0] ; char < []byte("~")[0]; char++{
		fmt.Printf("key=%s ", string(char))
		for i := 0; i < len(xored_bytes); i++{
			decrypted[i] = xored_bytes[i] ^ char
		}
		fmt.Println(string(decrypted))
	}
}
