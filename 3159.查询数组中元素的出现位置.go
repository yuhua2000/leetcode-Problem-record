func occurrencesOfElement(nums []int, queries []int, x int) []int {
	index := make([]int, 0, len(nums)/2)
	for i, num := range nums {
		if num == x {
			index = append(index, i)
		}
	}

	result := make([]int, len(queries))
	for i, q := range queries {
		if q > len(index) {
			result[i] = -1
		} else {
			result[i] = index[q-1]
		}
	}

	return result
}
