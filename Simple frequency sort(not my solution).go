package main

import "sort"

func Solve(arr []int) []int {
	m := map[int]int{}

	for _, d := range arr {
		m[d]++
	}

	sort.Slice(arr, func(i, j int) bool {
		if m[arr[i]] == m[arr[j]] {
			return arr[i] < arr[j]
		}
		return m[arr[i]] > m[arr[j]]
	})

	return arr
}
