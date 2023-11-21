package main

import (
	"strings"
)

func High(s string) string {
	a := make(map[rune]int, 0)
	for sym, point := 'a', 0; sym <= 'z'; sym++ {
		point++
		a[sym] = point
	}

	checkString := strings.Split(s, " ")
	score := make([]int, len(checkString))

	for idxWord, word := range checkString {
		sum := 0
		for _, symb := range word {
			sum += a[symb]
		}
		score[idxWord] = sum
	}

	max := score[0]
	result := 0
	for index, value := range score {
		if value > max {
			max = value
			result = index
		}
	}

	return checkString[result]
}
