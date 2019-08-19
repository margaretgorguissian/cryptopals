package main

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"sort"
)

// characterFrequency map from ricpacca/cryptopals

var characterFrequency = map[string]float64 {
    "a": 0.0651738, "b": 0.0124248, "c": 0.0217339, "d": 0.0349835, "e": 0.1041442, "f": 0.0197881, "g": 0.0158610,
    "h": 0.0492888, "i": 0.0558094, "j": 0.0009033, "k": 0.0050529, "l": 0.0331490, "m": 0.0202124, "n": 0.0564513,
    "o": 0.0596302, "p": 0.0137645, "q": 0.0008606, "r": 0.0497563, "s": 0.0515760, "t": 0.0729357, "u": 0.0225134,
    "v": 0.0082903, "w": 0.0171272, "x": 0.0013692, "y": 0.0145984, "z": 0.0007836, " ": 0.1918182,
}

var isAlpha, _ = regexp.Compile(`^[\w\s[:punct:] ]+$`)

func getEnglishScore(s string) float64{
	s = strings.ToLower(s)
	length := len(s)
	score := float64(0)
	if isAlpha.MatchString(s){
		for i := 0; i < length; i++{
			score = score + characterFrequency[string(s[i])]
		}
	}
	return score
}

func main() {
	xoredString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	xoredBytes, _ := hex.DecodeString(xoredString)
	decrypted := make([]byte, len(xoredBytes))
  
  scores := make([]float64, 10)
	scoreWordMap := make(map[float64]string)
	
  for char := []byte("!")[0] ; char < []byte("~")[0]; char++{
		for i := 0; i < len(xoredBytes); i++{
			decrypted_char := xoredBytes[i] ^ char
			decrypted[i] = decrypted_char
		}
		score := getEnglishScore(string(decrypted))
		scoreWordMap[score] = string(decrypted)
		scores = append(scores, score)
	}
	sort.Float64s(scores)
	fmt.Print(scoreWordMap[scores[len(scores) - 1]])
}
