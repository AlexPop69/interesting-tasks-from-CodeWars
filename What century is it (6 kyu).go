package main

import (
	"fmt"
	"strconv"
)

func WhatCentury(year string) string {
	cent, _ := strconv.Atoi(year[:1])
	fmt.Println(cent)
	cent++
	fmt.Println(cent)
	result := strconv.Itoa(cent)
	endings := []string{"st", "nd", "rd", "th"}

	a := cent % 10
	b := cent % 100
	switch {
	case (cent == 1 || cent >= 21) && (a == 1 || b == 1) && (b != 11):
		return result + endings[0]
	case cent == 2 || (a == 2 && cent > 20):
		return result + endings[1]
	case cent == 3 || (a == 3 && cent > 20):
		return result + endings[2]
	// case (cent >= 2 && cent <= 4 || cent > 21) && ((a >= 2 && a <= 4) && (b >= 2 && b <= 4)):
	// 	return fmt.Sprintf(endings[1])
	default:
		return result + endings[3]
	}
}
