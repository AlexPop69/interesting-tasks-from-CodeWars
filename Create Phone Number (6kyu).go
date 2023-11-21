package main

import "strconv"

func CreatePhoneNumber(numbers [10]uint) string {
	res := "("
	for i := 0; i < len(numbers); i++ {
		num := strconv.Itoa(int(numbers[i]))
		switch i {
		case 3:
			res += ") "
		case 6:
			res += "-"
		}
		res += num
	}
	return res
}
