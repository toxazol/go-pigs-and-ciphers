package gopigtranslator

import (
	"strings"
	"unicode"
)

func translateWord(word string) string {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true}

	prefix := make([]rune, len(word))
	suffix := make([]rune, len(word))
	postfix := "ay"
	if vowels[rune(word[0])] {
		postfix = "yay"
	}
	var prefixEnded bool

	for _, symbol := range word {
		if !prefixEnded && !vowels[unicode.ToLower(symbol)] {
			prefix = append(prefix, symbol)
		} else {
			prefixEnded = true
			suffix = append(suffix, symbol)
		}

	}
	return string(suffix) + string(prefix) + postfix
}

// Translate translates your english text to Pig Latin
func Translate(englishText string) string {
	result := make([]string, 0)
	for _, word := range strings.Split(englishText, " ") {
		result = append(result, translateWord(word))
	}
	return strings.Join(result, " ")
}
