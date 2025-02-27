package main

import (
	"strings"
)

const StemEndings = "аяеєиіїоуюь"

// Recursively extracts the stem of a word
func GetStem(word []rune) []rune {
	if l := len(word); l > 2 && strings.ContainsRune(StemEndings, word[l-1]) {
		word = GetStem(word[:l-1])
	}

	return word
}

func getLastLetter(word string) string {
	runes := []rune(word)
	return string(runes[len(runes)-1])
}
