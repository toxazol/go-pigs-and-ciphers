package gocipher

import (
	"unicode"
)

// Cipher is an object that can help you secretly communicate with your friend
type Cipher struct {
	EncodeMap map[rune]rune
	DecodeMap map[rune]rune
}

func (cr Cipher) code(in string, cipherMap map[rune]rune) string {
	result := make([]rune, len(in))
	for _, symbol := range in {
		toSymbol, ok := cipherMap[unicode.ToLower(symbol)]
		if ok {
			result = append(result, toSymbol)
		} else {
			result = append(result, symbol)
		}

	}
	return string(result)
}

// Encode your message before send
func (cr Cipher) Encode(in string) string {
	return cr.code(in, cr.EncodeMap)
}

// Decode the reply
func (cr Cipher) Decode(in string) string {
	if cr.DecodeMap == nil {
		cr.DecodeMap = make(map[rune]rune)
		for k, v := range cr.EncodeMap {
			cr.DecodeMap[v] = k
		}
	}
	return cr.code(in, cr.DecodeMap)
}
