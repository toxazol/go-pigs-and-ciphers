package main

import (
	"fmt"
	"os"

	"github.com/toxazol/go-pigs-and-ciphers/gocipher"
	"github.com/toxazol/go-pigs-and-ciphers/gopigtranslator"
)

func main() {
	var translator, method, text string
	if len(os.Args) == 4 {
		translator = os.Args[1]
		method = os.Args[2]
		text = os.Args[3]
	} else {
		fmt.Println("wrong number of arguments. usage: go run main.go translator text")
		return
	}

	if translator == "piglatin" {
		fmt.Println(gopigtranslator.Translate(text))
	} else if translator == "cipher" {
		numberCipher := gocipher.Cipher{
			EncodeMap: map[rune]rune{
				'a': '1',
				'e': '2',
				'i': '3',
				'o': '4',
				'u': '5',
			}}
		if method == "encode" {
			fmt.Println(numberCipher.Encode(text))
		} else if method == "decode" {
			fmt.Println(numberCipher.Decode(text))
		} else {
			fmt.Println("unsupported method. Choose encode or decode")
		}
	} else {
		fmt.Println("unsupported translator. Choose piglatin or cipher")
	}
}
