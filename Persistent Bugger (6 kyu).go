package main

import (
	"strconv"
	"strings"
)

func Persistence(n int) int {
	count := 0

	for n > 9 {
		number := strconv.Itoa(n)
		numbers := strings.Split(number, "")

		for i, j := 0, 1; i < len(numbers); i++ {
			num, _ := strconv.Atoi(numbers[i])
			j *= num
			n = j
		}
		count++
	}
	return count
}
