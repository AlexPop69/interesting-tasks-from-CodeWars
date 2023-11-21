package main

//There are no explanations. You have to create the code that gives the following results in Go:
// oneTwoThree(0) == ['0', '0']
// oneTwoThree(1) == ['1', '1']
// oneTwoThree(3) == ['3', '111']
// oneTwoThree(19) == ['991', '1111111111111111111']

import (
	"strconv"
	"strings"
)

func OneTwoThreeSol(n int) [2]string {
	if n != 0 {
		firstString := ""
		secondString := ""

		if n < 9 {
			firstString = strconv.Itoa(n)
		}
		if n >= 9 {
			firstString = strings.Repeat("9", n/9)
			if n%9 != 0 {
				firstString += strconv.Itoa(n % 9)
			}
		}
		secondString = strings.Repeat("1", n)
		return [2]string{firstString, secondString}
	}

	return [2]string{"0", "0"}
}
