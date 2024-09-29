package main

import (
	"strconv"
)

func reverseFunc(s string) (result string) {
	var reverseString string

	for _, v := range s {
		_, err := strconv.Atoi(string(v))
		if err == nil {
			reverseString = reverseString + string(v)
		} else {
			reverseString = string(v) + reverseString
		}

	}

	return reverseString

}
