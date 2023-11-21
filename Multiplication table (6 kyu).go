package main

import "fmt"

//Your task, is to create N×N multiplication table, of size provided in parameter.
//For example, when given size is 3:
//	1 2 3
//	2 4 6
//	3 6 9
//For the given example, the return value should be:
//[[1,2,3],[2,4,6],[3,6,9]]

func MultiplicationTable(size int) [][]int {
	m := make([][]int, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			m[i] = append(m[i], (j+1)*(i+1))
		}
	}
	//output such as a table
	for _, v1 := range m {
		fmt.Println(v1)
	}
	return m
}
