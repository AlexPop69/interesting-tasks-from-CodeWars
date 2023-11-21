package main

import (
	"strconv"
	"strings"
)

//Definition
//Strong number is the number that the sum of the factorial of its digits is equal to number itself.
//For example, 145 is strong, since 1! + 4! + 5! = 1 + 24 + 120 = 145.

//Given a number, Find if it is Strong or not and return either "STRONG!!!!" or "Not Strong !!".

// Notes:
// Number passed is always Positive.
// Return the result as String

func Strong(n int) string {
	str := strings.Split(strconv.Itoa(n), "")
	res := 0

	for _, v := range str {
		sum := 1
		num, _ := strconv.Atoi(v)
		for i := 1; i <= num; i++ {
			sum *= i
		}
		res += sum
	}
	if n == res {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}
