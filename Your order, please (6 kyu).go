package main

import (
	"strconv"
	"strings"
	"unicode"
)

func Order(sentence string) string {
	//split sentence on the words
	str := strings.Split(sentence, " ")
	//make empty slice of strings for return
	result := make([]string, len(str))
	//iterate for words
	for _, word := range str {
		//iterate for symbols in the word
		for i := 0; i < len(word); i++ {
			// if sybmbol is number
			if unicode.IsDigit(rune(word[i])) {
				//covert that symbol to integer and write to position
				position, _ := strconv.Atoi(string(word[i]))
				//write word to position on the result
				result[position-1] = word
			}
		}
	}
	return strings.Join(result, " ")
}
