/*
3/28/2023
ECB
*/
package main

import (
	"fmt"
	"strconv"
)

/* an array with 4 rows and 2 columns*/
var codebook = [2][4][2]int{{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}, {{0b00, 0b10}, {0b01, 0b11}, {0b10, 0b00}, {0b11, 0b01}}}
var message = [4]int{0b00, 0b01, 0b10, 0b11}
var key int = 0

func textToBinary(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%.8b", binString, c)
	}
	return binString
}
func binaryToString(binaryString string) (textString string) {
	textString = ""
	var start, end int = 0, 8
	for end < len(binaryString)+1 {
		binaryInt, _ := strconv.ParseInt(binaryString[start:end], 2, 0)
		text := string(binaryInt)
		textString += text
		end += 8
		start += 8
	}
	return textString
}

func codebookLookup(key int) (ciphertext int) {
	var i, j, k int = 0, 0, 1
	var found bool = false
	for i = 0; i < 2; i++ {
		for j = 0; j < 4; j++ {
			if codebook[i][j][k] == key {
				k--
				ciphertext = codebook[i][j][k]
				found = true
				break
			}
		}
		if found {
			break
		}

	}
	return ciphertext
}

func main() {
	var text string

	fmt.Printf("Enter your message: ")
	fmt.Scanln(&text)
	text = textToBinary(text)
	fmt.Println("Binary plaintext: " + text)

	var x, i int = 0, 0
	var lookupValue int = 0
	var ciphertext string = ""
	var plaintextBinary string = ""
	for x = 0; x < len(text)-1; x += 2 {
		binaryInt, _ := strconv.ParseInt(text[x:x+2], 2, 0)
		t := strconv.FormatInt(binaryInt, 10)
		for i = 0; i < len(codebook[key]); i++ {
			//println(strconv.Itoa(codebook[key][i][1]))
			if t == strconv.Itoa(codebook[key][i][1]) {
				lookupValue = codebook[key][i][0]
				fmt.Printf("The ciphered value is %b\n", lookupValue)
				ciphertext += fmt.Sprintf("%02b", lookupValue)
				break
			}

		}
	}
	println("Ciphertext: " + ciphertext)

	var k, l int = 0, 0
	for k = 0; k < len(ciphertext)-1; k += 2 {
		binaryInt, _ := strconv.ParseInt(ciphertext[k:k+2], 2, 0)
		t := strconv.FormatInt(binaryInt, 10)
		for l = 0; l < len(codebook[key]); l++ {
			if t == strconv.Itoa(codebook[key][l][0]) {
				lookupValue = codebook[key][l][1]
				fmt.Printf("The deciphered value is %b\n", lookupValue)
				plaintextBinary += fmt.Sprintf("%02b", lookupValue)
				break
			}
		}
	}
	println("deciphered plaintext: " + plaintextBinary)
	plaintext := binaryToString(plaintextBinary)
	println("deciphered text: " + plaintext)
}
