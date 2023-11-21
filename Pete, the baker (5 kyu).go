package main

func Cakes(recipe, available map[string]int) int {
	res := make([]int, 0)
	for ingredient := range recipe {
		res = append(res, available[ingredient]/recipe[ingredient])
	}

	min := res[0]
	for _, v := range res {
		if v <= min {
			min = v
		}
	}
	return min
}
