/*
3/28/2023
ECB
*/
package main

import (
	"fmt"
	"strconv"
)

var codebook = [2][4][2]int{{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}},
	{{0b00, 0b10}, {0b01, 0b11}, {0b10, 0b00}, {0b11, 0b01}}}
var message = [4]int{0b00, 0b01, 0b10, 0b11}
var key int = 0 // Uses codebook1, the first array.
var lookupValue int = 0

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

func ECB_Ciphertext(text string) (ciphertext string) {
	ciphertext = ""
	var x, i int = 0, 0
	for x = 0; x < len(text)-1; x += 2 {
		binaryInt, _ := strconv.ParseInt(text[x:x+2], 2, 0)
		t := strconv.FormatInt(binaryInt, 10)
		for i = 0; i < len(codebook[key]); i++ {
			if t == strconv.Itoa(codebook[key][i][1]) {
				lookupValue = codebook[key][i][0]
				fmt.Printf("The ciphered value is %b\n", lookupValue)
				ciphertext += fmt.Sprintf("%02b", lookupValue)
				break
			}

		}
	}
	println("Ciphertext: " + ciphertext)
	return ciphertext
}

func ECB_Decipher(ciphertext string) (plaintext string) {
	var plaintextBinary string = ""
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
	plaintext = binaryToString(plaintextBinary)
	println("deciphered text: " + plaintext)
	return plaintext
}

func main() {
	var text string
	fmt.Printf("Enter your message: ")
	fmt.Scanln(&text)
	text = textToBinary(text)
	fmt.Println("Binary plaintext: " + text)
	ciphertext := ECB_Ciphertext(text)
	ECB_Decipher(ciphertext)
}

