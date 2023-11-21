package main

func FindUniq(arr []float32) float32 {
	//make empty map
	result := make(map[float32]int)
	//iterate for arr
	for i := 0; i < len(arr); i++ {
		// add elements to map and every elements such as "[arr[i]] quantity of elements"
		result[arr[i]] += 1
	}
	//iterate for map
	for k, v := range result {
		//if quantity of elements equals one
		if v == 1 {
			return k
		}
	}
	return 0
}
