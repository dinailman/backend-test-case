package main

import "strings"

func longestString(s string) (result string, length int) {
	best, length := "", 0
	for _, word := range strings.Split(s, " ") {
		if len(word) > length {
			best, length = word, len(word)
		}
	}
	return best, length
}
