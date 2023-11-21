package main

//Write an algorithm that takes an array and moves all of the zeros to the end, preserving the order of the other elements.
//MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }

func MoveZeros(arr []int) []int {
	arrZero := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			arrZero = append(arrZero, arr[i])
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	result := append(arr, arrZero...)
	return result
}

// func MoveZeros(arr []int) []int {
// 	result := make([]int, len(arr))
// 	idx := 0

// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != 0 {
// 			result[idx] = arr[i]
// 			idx++
// 		}
// 	}
// 	return result
// }
