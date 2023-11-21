package main

//найти число встречающееся нечетноe количество раз
func FindOdd(seq []int) int {
	odd := make(map[int]int, 0)
	for _, v := range seq {
		odd[v]++
	}

	for v, k := range odd {
		if k%2 != 0 {
			return v
		}
	}
	return 0
}
