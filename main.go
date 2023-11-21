package main

import (
	"fmt"
)

func main() {
	//Find the unique number
	s := []float32{1.0, 1, 1, 2, 1, 1}
	fmt.Println(FindUniq(s))

	//Highest Scoring Word
	sec := "man i need a taxi up to ubud"
	fmt.Println(High(sec))

	//What a Century?
	WhatCentury("2523")

	//Buying a car
	fmt.Println(NbMonths(7500, 7500, 1000, 1.55))

	//Playing with digits
	fmt.Println(DigPow(46288, 3)) //51

	//Human Readable Time
	fmt.Println(HumanReadableTime(359999))

	//Tortoise racing
	fmt.Println(Race(720, 850, 70))

	//Strong Number (Special Numbers Series #2)
	fmt.Println(Strong(145)) //strong

	//Moving Zeros To The End
	fmt.Println(MoveZeros([]int{0, 1, 0, 3, 12}))

	//Find the odd int
	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))

	//Begin your day with a challenge, but an easy one.
	fmt.Println(OneTwoThreeSol(19)) // [991 111111111111111111]

	//Multiplication table
	sizeOfTable := 3
	fmt.Println(MultiplicationTable(sizeOfTable))

	//Pete, the baker
	m1 := map[string]int{"flour": 500, "sugar": 200, "eggs": 1}
	m2 := map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}
	fmt.Println(Cakes(m1, m2))

	//Simple frequency sort
	fmt.Println(Solve([]int{2, 3, 5, 3, 7, 9, 5, 3, 7})) //3,3,3,5,5,7,7,2,9

	//Create Phone Number
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))

	//Persistent Bugger
	fmt.Println(Persistence(999)) //4

	//Your order, please
	fmt.Println(Order("4of Fo1r pe6ople g3ood th5e the2")) //Fo1r the2 g3ood 4of th5e pe6ople
}
