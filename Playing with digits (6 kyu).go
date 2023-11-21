package main

import (
	"math"
	"strconv"
	"strings"
)

func DigPow(n, p int) int {
	nStr := strconv.Itoa(n)
	numbersStr := strings.Split(nStr, "")

	sum := 0
	for i := 0; i < len(numbersStr); i++ {
		num, _ := strconv.Atoi(numbersStr[i])
		sum += int(math.Pow(float64(num), float64(p)))
		p++
	}

	if sum%n == 0 {
		return sum / n
	}
	return -1
}
