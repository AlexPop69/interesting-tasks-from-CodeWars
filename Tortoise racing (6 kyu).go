package main

// func Race(v1, v2, g int) [3]int {
// 	if v1 < v2 {
// 		var speed1 float64 = float64(v1) / 3600
// 		var speed2 float64 = float64(v2) / 3600
// 		tor1 := float64(g)
// 		tor2 := 0.0
// 		time := 0

// 		for {
// 			tor1 += speed1
// 			tor2 += speed2

// 			if tor2 >= tor1 {
// 				break
// 			}
// 			time++
// 		}

// 		return [3]int{
// 			time / 3600,
// 			time % 3600 / 60,
// 			time % 3600 % 60,
// 		}
// 	}
// 	return [3]int{-1, -1, -1}
// }

func Race(v1, v2, g int) [3]int {
	if v1 < v2 {
		time := g * 3600 / (v2 - v1)
		return [3]int{
			time / 3600,
			time % 3600 / 60,
			time % 3600 % 60,
		}
	}
	return [3]int{-1, -1, -1}
}
